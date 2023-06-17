package devicestore

func GetUserIdFilter(userId string) map[string]interface{} {
	return map[string]interface{}{
		"user_id": userId,
	}
}
