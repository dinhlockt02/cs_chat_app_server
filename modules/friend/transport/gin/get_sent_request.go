package friendgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendbiz "cs_chat_app_server/modules/friend/biz"
	friendstore "cs_chat_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetSentRequest(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		senderId := requester.GetId()

		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		getSentRequestBiz := friendbiz.NewGetSentRequestBiz(friendStore)
		result, err := getSentRequestBiz.GetSentRequest(context.Request.Context(), senderId)
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
