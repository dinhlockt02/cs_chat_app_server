package appcontext

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type AppContext interface {
	MongoClient() *mongo.Client
}

type appContext struct {
	mongoClient *mongo.Client
}

func NewAppContext(
	mongoClient *mongo.Client,
) *appContext {
	return &appContext{
		mongoClient: mongoClient,
	}
}

func (a *appContext) MongoClient() *mongo.Client {
	return a.mongoClient
}
