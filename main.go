package main

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"os"
	"time"
)

func init() {

	setupLogger()

	var err error

	common.AppDatabase = os.Getenv("MONGO_DB")

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

	// Create app context

	appCtx := appcontext.NewAppContext(client)

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
