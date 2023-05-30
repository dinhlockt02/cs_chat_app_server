package common

import "go.mongodb.org/mongo-driver/bson/primitive"

// AddIdFilter is an method that will add id filter to provied filter
//
// It required id as string, and has the mongoid format.
//
// It will return [ErrInvalidRequest] if the id is not has the right format
func AddIdFilter(filter map[string]interface{}, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidRequest(err)
	}
	filter["_id"] = _id
	return nil
}

func GetOrFilter(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"$or": filters,
	}
}

func GetAndFilter(filters ...map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"$and": filters,
	}
}

func GetExistsFilter(fieldName string, exists bool) map[string]interface{} {
	return map[string]interface{}{
		fieldName: map[string]interface{}{
			"$exists": exists,
		},
	}
}

func GetInFilter(fieldName string, values ...interface{}) map[string]interface{} {
	return map[string]interface{}{
		fieldName: map[string]interface{}{
			"$in": values,
		},
	}
}