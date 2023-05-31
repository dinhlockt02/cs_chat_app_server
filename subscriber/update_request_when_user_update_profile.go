package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
	userstore "cs_chat_app_server/modules/user/store"
)

func UpdateRequestWhenUserUpdateProfile(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)

	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
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
					panic(err)
				}
				if user == nil {
					return
				}

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestSenderIdFilter(*user.Id),
					&requestmdl.UpdateRequest{
						Sender: &requestmdl.RequestUser{
							Id:     *user.Id,
							Name:   user.Name,
							Avatar: user.Avatar,
						},
					})
				if err != nil {
					return
				}

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestReceiverIdFilter(*user.Id),
					&requestmdl.UpdateRequest{
						Receiver: &requestmdl.RequestUser{
							Id:     *user.Id,
							Name:   user.Name,
							Avatar: user.Avatar,
						},
					})
				if err != nil {
					return
				}

			}(userId)
		}
	}()
}
