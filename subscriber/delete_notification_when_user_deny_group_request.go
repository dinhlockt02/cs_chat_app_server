package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	notimodel "cs_chat_app_server/components/notification/model"
	notistore "cs_chat_app_server/components/notification/store"
	"github.com/rs/zerolog/log"
)

func DeleteNotificationWhenUserDenyGroupRequest(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicDenyGroupRequest)

	notificationStore := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	go func() {

		for reqId := range ch {
			go func(rid string) {
				defer common.Recovery()
				err := notificationStore.Delete(ctx, notistore.GetDirectFilter(rid, notimodel.Request))
				if err != nil {
					log.Error().Err(err).Str("package", "subscriber.DeleteNotificationWhenUserAcceptFriendRequest").Send()
					panic(err)
				}
			}(reqId)
		}
	}()
}
