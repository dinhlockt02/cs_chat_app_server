package appcontext

import (
	fbs "cs_chat_app_server/components/firebase"
	"cs_chat_app_server/components/hasher"
	"cs_chat_app_server/components/mailer"
	notirepo "cs_chat_app_server/components/notification/repository"
	"cs_chat_app_server/components/pubsub"
	"cs_chat_app_server/components/socket"
	"cs_chat_app_server/components/tokenprovider"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	MongoClient() *mongo.Client
	TokenProvider() tokenprovider.TokenProvider
	Hasher() hasher.Hasher
	Mailer() mailer.Mailer
	RedisClient() *redis.Client
	FirebaseApp() fbs.FirebaseApp
	Socket() socket.Socket
	Notification() notirepo.NotificationServiceRepository
	PubSub() pubsub.PubSub
}

type appContext struct {
	mongoClient   *mongo.Client
	tokenProvider tokenprovider.TokenProvider
	hasher        hasher.Hasher
	mailer        mailer.Mailer
	redisClient   *redis.Client
	fa            fbs.FirebaseApp
	socket        socket.Socket
	notification  notirepo.NotificationServiceRepository
	pubsub        pubsub.PubSub
}

func NewAppContext(
	mongoClient *mongo.Client,
	tokenProvider tokenprovider.TokenProvider,
	hasher hasher.Hasher,
	mailer mailer.Mailer,
	redisClient *redis.Client,
	fa fbs.FirebaseApp,
	socket socket.Socket,
	notification notirepo.NotificationServiceRepository,
	pubsub pubsub.PubSub,
) *appContext {
	return &appContext{
		mongoClient:   mongoClient,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		mailer:        mailer,
		redisClient:   redisClient,
		fa:            fa,
		socket:        socket,
		notification:  notification,
		pubsub:        pubsub,
	}
}

func (a *appContext) MongoClient() *mongo.Client {
	return a.mongoClient
}

func (a *appContext) TokenProvider() tokenprovider.TokenProvider {
	return a.tokenProvider
}

func (a *appContext) Hasher() hasher.Hasher {
	return a.hasher
}

func (a *appContext) Mailer() mailer.Mailer {
	return a.mailer
}

func (a *appContext) RedisClient() *redis.Client {
	return a.redisClient
}

func (a *appContext) FirebaseApp() fbs.FirebaseApp {
	return a.fa
}

func (a *appContext) Socket() socket.Socket {
	return a.socket
}

func (a *appContext) Notification() notirepo.NotificationServiceRepository {
	return a.notification
}

func (a *appContext) PubSub() pubsub.PubSub {
	return a.pubsub
}
