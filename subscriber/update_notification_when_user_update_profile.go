package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	notimodel "cs_chat_app_server/components/notification/model"
	notistore "cs_chat_app_server/components/notification/store"
	usermodel "cs_chat_app_server/modules/user/model"
	userstore "cs_chat_app_server/modules/user/store"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

func UpdateNotificationWhenUserUpdateProfile(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)

	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for userId := range ch {
			go func(uid string) {
				defer common.Recovery()

				filter := make(map[string]interface{})

				err := common.AddIdFilter(filter, uid)
				if err != nil {
					panic(err)
				}

				user, err := userStore.Find(context.Background(), filter)
				if err != nil {
					log.Error().Err(err)
					return
				}
				if user == nil {
					log.Error().Err(errors.New(fmt.Sprintf("user %s not found", uid)))
					return
				}

				// Update subject
				go func(usr *usermodel.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetSubjectFilter(*usr.Id),
						&notimodel.UpdateNotification{Subject: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    *usr.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update direct
				go func(usr *usermodel.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetDirectFilter(*usr.Id),
						&notimodel.UpdateNotification{Direct: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    *usr.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update indirect
				go func(usr *usermodel.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetIndirectFilter(*usr.Id),
						&notimodel.UpdateNotification{Indirect: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    *usr.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

				// Update prep
				go func(usr *usermodel.User) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetPrepFilter(*usr.Id),
						&notimodel.UpdateNotification{Prep: &notimodel.NotificationObject{
							Name:  user.Name,
							Image: &user.Avatar,
							Id:    *usr.Id,
							Type:  notimodel.User,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(user)

			}(userId)
		}
	}()
}
