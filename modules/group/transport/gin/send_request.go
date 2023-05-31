package groupgin

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupbiz "cs_chat_app_server/modules/group/biz"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	requeststore "cs_chat_app_server/modules/request/store"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func SendGroupRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		requester := c.MustGet(common.CurrentUser).(common.Requester)
		friendId := c.Param("friendId")
		groupId := c.Param("groupId")

		if !primitive.IsValidObjectID(friendId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}
		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)

		getGroupBiz := groupbiz.NewGetGroupBiz(groupRepo)

		group, err := getGroupBiz.GetById(context.Background(), groupId)
		if err != nil {
			panic(err)
		}

		sendGroupRequestBiz := groupbiz.NewSendGroupRequestBiz(groupRepo)
		err = sendGroupRequestBiz.SendRequest(context.Background(), requester.GetId(), friendId, group)
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusCreated, gin.H{"data": true})
	}
}
