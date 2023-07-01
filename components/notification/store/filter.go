package notistore

import notimodel "cs_chat_app_server/components/notification/model"

func GetSubjectFilter(id string, typ notimodel.NotificationObjectType) map[string]interface{} {
	return map[string]interface{}{
		"subject.id":   id,
		"subject.type": typ,
	}
}

func GetDirectFilter(id string, typ notimodel.NotificationObjectType) map[string]interface{} {
	return map[string]interface{}{
		"direct.id":    id,
		"subject.type": typ,
	}
}

func GetIndirectFilter(id string, typ notimodel.NotificationObjectType) map[string]interface{} {
	return map[string]interface{}{
		"indirect.id":  id,
		"subject.type": typ,
	}
}

func GetPrepFilter(id string, typ notimodel.NotificationObjectType) map[string]interface{} {
	return map[string]interface{}{
		"prep.id":      id,
		"subject.type": typ,
	}
}

func GetOwnerFilter(id string) map[string]interface{} {
	return map[string]interface{}{
		"owner": id,
	}
}
