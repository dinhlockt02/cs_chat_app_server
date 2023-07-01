package friendgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendbiz "cs_chat_app_server/modules/friend/biz"
	friendstore "cs_chat_app_server/modules/friend/store"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	requeststore "cs_chat_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Unfriend(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		userId := requester.GetId()
		friendId := context.Param("id")
		if !primitive.IsValidObjectID(friendId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(groupStore, requestStore)
		unfriendBiz := friendbiz.NewUnfriendBiz(friendStore, groupRepo, appCtx.PubSub())
		if err := unfriendBiz.Unfriend(context.Request.Context(), userId, friendId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
