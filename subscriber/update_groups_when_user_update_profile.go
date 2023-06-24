package subscriber

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupmdl "cs_chat_app_server/modules/group/model"
	groupstore "cs_chat_app_server/modules/group/store"
	userstore "cs_chat_app_server/modules/user/store"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
)

func UpdateGroupsWhenUserUpdateProfile(appCtx appcontext.AppContext, ctx context.Context) {
	ch := appCtx.PubSub().Subscribe(ctx, common.TopicUserUpdateProfile)

	userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
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

				err = groupStore.UpdateGroupMember(
					context.Background(),
					groupstore.GetMemberIdInGroupMembersFilter(uid),
					&groupmdl.GroupUser{
						Id:     uid,
						Name:   user.Name,
						Avatar: user.Avatar,
					})
				if err != nil {
					log.Error().Err(err).Msg("")
					return
				}

			}(userId)
		}
	}()
}
