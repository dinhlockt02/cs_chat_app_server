package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoId struct {
	Id *string `bson:"_id" json:"id"`
}

type MongoUpdatedAt struct {
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

func ToObjectId(hex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hex)
}
