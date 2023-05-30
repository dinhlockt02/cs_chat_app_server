package notiservice

import (
	"context"
	"cs_chat_app_server/common"
	notimodel "cs_chat_app_server/components/notification/model"
	"encoding/json"
	"firebase.google.com/go/v4/messaging"
	"github.com/rs/zerolog/log"
	"time"
)

type firebaseNotificationService struct {
	client *messaging.Client
}

func NewFirebaseNotificationService(client *messaging.Client) firebaseNotificationService {
	return firebaseNotificationService{
		client: client,
	}
}

func (service firebaseNotificationService) Push(ctx context.Context, token []string, notification *notimodel.Notification) error {
	title, body := notification.GetMessage()

	content, err := notification.GetContent()
	if err != nil {
		return common.ErrInternal(err)
	}

	marshaledContent, err := json.Marshal(content)

	if err != nil {
		return common.ErrInternal(err)
	}

	marshaledActionButton, err := json.Marshal(notification.GetActionButton())
	if err != nil {
		return common.ErrInternal(err)
	}

	zeroDuration := time.Duration(0)

	msg := &messaging.MulticastMessage{
		Tokens: token,
		Android: &messaging.AndroidConfig{
			Priority: "high",
			TTL:      &zeroDuration,
			//Notification: &messaging.AndroidNotification{
			//	Title:    title,
			//	Body:     body,
			//	ImageURL: *notification.Prep.Image,
			//},
		},
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Data: map[string]string{
			"content":       string(marshaledContent),
			"actionButtons": string(marshaledActionButton),
		},
	}

	br, err := service.client.SendMulticast(ctx, msg)
	if err != nil {
		return common.ErrInternal(err)
	}
	log.Info().Msgf("%v sent successfully", br.SuccessCount)
	log.Info().Msgf("%v sent failed", br.FailureCount)
	return nil
}
