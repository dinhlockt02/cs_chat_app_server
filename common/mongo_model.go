package common

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type MongoId struct {
	Id *string `bson:"_id,omitempty" json:"id,omitempty"`
}

type MongoUpdatedAt struct {
	UpdatedAt *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}

type MongoCreatedAt struct {
	CreatedAt *time.Time `bson:"created_at" json:"created_at,omitempty"`
}

func ToObjectId(hex string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(hex)
}

func GetTextSearch(value string, caseSensitive bool, diacriticSensitive bool) map[string]interface{} {
	return map[string]interface{}{
		"$text": map[string]interface{}{
			"$search":             value,
			"$caseSensitive":      caseSensitive,
			"$diacriticSensitive": diacriticSensitive,
		},
	}
}
