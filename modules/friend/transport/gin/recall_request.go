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

func RecallRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		senderId := requester.GetId()
		receiverId := context.Param("id")

		if !primitive.IsValidObjectID(senderId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}
		if !primitive.IsValidObjectID(receiverId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		recallRequestBiz := friendbiz.NewRecallRequestBiz(friendStore)
		if err := recallRequestBiz.RecallRequest(context.Request.Context(), senderId, receiverId); err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}