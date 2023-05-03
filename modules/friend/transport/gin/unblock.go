package friendgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendbiz "cs_chat_app_server/modules/friend/biz"
	friendstore "cs_chat_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Unblock(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		userId := requester.GetId()
		blockedId := context.Param("id")
		if !primitive.IsValidObjectID(blockedId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		unblockBiz := friendbiz.NewUnblockBiz(friendStore)
		if err := unblockBiz.Unblock(context.Request.Context(), userId, blockedId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
