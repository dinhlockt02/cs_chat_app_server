package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	notimodel "cs_chat_app_server/components/notification/model"
	notistore "cs_chat_app_server/components/notification/store"
	groupmdl "cs_chat_app_server/modules/group/model"
	groupstore "cs_chat_app_server/modules/group/store"
	"github.com/rs/zerolog/log"
)

func UpdateNotificationWhenGroupUpdated(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicGroupUpdated)

	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for groupId := range ch {
			go func(gid string) {
				defer common.Recovery()

				filter := make(map[string]interface{})

				err := common.AddIdFilter(filter, gid)
				if err != nil {
					panic(err)
				}

				group, err := groupStore.FindGroup(context.Background(), filter)
				if err != nil {
					panic(err)
				}
				if group == nil {
					return
				}

				// Update subject
				go func(grp *groupmdl.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetSubjectFilter(*grp.Id, notimodel.Group),
						&notimodel.UpdateNotification{Subject: &notimodel.NotificationObject{
							Name:  grp.Name,
							Image: grp.ImageUrl,
							Id:    *grp.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update direct
				go func(grp *groupmdl.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetDirectFilter(*grp.Id, notimodel.Group),
						&notimodel.UpdateNotification{Direct: &notimodel.NotificationObject{
							Name:  grp.Name,
							Image: grp.ImageUrl,
							Id:    *grp.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update indirect
				go func(grp *groupmdl.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetIndirectFilter(*grp.Id, notimodel.Group),
						&notimodel.UpdateNotification{Indirect: &notimodel.NotificationObject{
							Name:  grp.Name,
							Image: grp.ImageUrl,
							Id:    *grp.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

				// Update prep
				go func(grp *groupmdl.Group) {
					err = notificationStore.UpdateNotifications(
						context.Background(),
						notistore.GetPrepFilter(*grp.Id, notimodel.Group),
						&notimodel.UpdateNotification{Prep: &notimodel.NotificationObject{
							Name:  grp.Name,
							Image: grp.ImageUrl,
							Id:    *grp.Id,
							Type:  notimodel.Group,
						}})
					if err != nil {
						log.Error().Err(err).Msg(err.Error())
					}
				}(group)

			}(groupId)
		}
	}()
}
