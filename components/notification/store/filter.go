package notistore

func GetSubjectFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"subject.id": id,
	}
}

func GetDirectFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"direct.id": id,
	}
}

func GetIndirectFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"indirect.id": id,
	}
}

func GetPrepFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"prep.id": id,
	}
}
