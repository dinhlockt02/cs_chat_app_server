package groupgin

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupbiz "cs_chat_app_server/modules/group/biz"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	gchatbiz "cs_chat_app_server/modules/group_chat/biz"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
	gchatstore "cs_chat_app_server/modules/group_chat/store"
	requeststore "cs_chat_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func AcceptRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		requesterId := requester.GetId()
		groupId := c.Param("groupId")

		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}
		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		acceptRequestBiz := groupbiz.NewAcceptGroupRequestBiz(groupRepo, appCtx.PubSub())
		if err := acceptRequestBiz.AcceptRequest(c.Request.Context(), requesterId, groupId); err != nil {
			panic(err)
		}

		go func() {
			defer common.Recovery()
			store := gchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
			repo := gchatrepo.NewGroupChatRepository(store)
			biz := gchatbiz.NewSendMessageBiz(repo, appCtx.PubSub())
			chatItem := &gchatmdl.GroupChatItem{
				Type:        gchatmdl.System,
				SenderId:    requester.GetId(),
				GroupId:     groupId,
				SystemEvent: common.GetPointer(gchatmdl.MemberJoined),
			}
			_ = chatItem.Process()
			err := biz.Send(context.Background(), chatItem)
			if err != nil {
				log.Error().Err(err).Str("package", "groupgin.AcceptRequest.Send").Send()
			}
		}()
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
