package main

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	fbs "cs_chat_app_server/components/firebase"
	"cs_chat_app_server/components/hasher"
	"cs_chat_app_server/components/mailer"
	"cs_chat_app_server/components/tokenprovider/jwt"
	"cs_chat_app_server/middleware"
	v1route "cs_chat_app_server/route/v1"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/api/option"
	"net/http"
	"os"
	"strconv"
	"time"
)

func init() {

	setupLogger()

	var err error

	common.AppDatabase = os.Getenv("MONGO_DB")
	common.AccessTokenExpiry, err = strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXPIRY"))

	if err != nil {
		log.Panic().Msg(err.Error())
	}
}

func main() {

	// Get mongo client
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := connectMongoDB(ctx)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	// Get token provider

	tokenProvider := jwt.NewJwtTokenProvider(os.Getenv("SECRET"))

	// Create bcrypt hasher

	bcryptHasher := hasher.NewBcryptHasher()

	// Create mailer

	sendgridMailer := mailer.NewSendGridMailer(
		os.Getenv("SENDGRID_SENDER_NAME"),
		os.Getenv("SENDGRID_SENDER_EMAIL"),
		os.Getenv("SENDGRID_API_KEY"),
	)

	// Create redis client
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// Create Firebase App
	opt := option.WithCredentialsFile("./service-account-key.json")
	fa, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Panic().Msg(err.Error())
	}
	app := fbs.NewFirebaseApp(fa)

	// Create app context

	appCtx := appcontext.NewAppContext(client, tokenProvider, bcryptHasher, sendgridMailer, redisClient, app)

	envport := os.Getenv("SERVER_PORT")
	if envport == "" {
		envport = "8080"
	}
	port := fmt.Sprintf(":%v", envport)

	r := gin.New()

	r.Use(gin.Recovery())

	r.Use(middleware.Recover(appCtx))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1route.InitRoute(r, appCtx)

	if err := r.Run(port); err != nil {
		log.Panic().Msg(err.Error())
	}
}

func connectMongoDB(ctx context.Context) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	log.Info().Msg("Connect to mongodb successful")
	return client, nil
}

func setupLogger() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if gin.Mode() == gin.DebugMode {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
