package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupstore "cs_chat_app_server/modules/group/store"
	requestmdl "cs_chat_app_server/modules/request/model"
	requeststore "cs_chat_app_server/modules/request/store"
)

func UpdateRequestWhenGroupUpdated(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicGroupUpdated)

	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
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

				group, err := groupStore.FindGroup(context.Background(), filter)
				if err != nil {
					panic(err)
				}
				if group == nil {
					return
				}

				err = requestStore.UpdateRequests(
					context.Background(),
					requeststore.GetRequestGroupIdFilter(*group.Id),
					&requestmdl.UpdateRequest{
						Group: &requestmdl.RequestGroup{
							Id:       *group.Id,
							Name:     group.Name,
							ImageUrl: *group.ImageUrl,
						},
					})
				if err != nil {
					return
				}

			}(userId)
		}
	}()
}
