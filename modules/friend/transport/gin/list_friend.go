package friendgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendbiz "cs_chat_app_server/modules/friend/biz"
	friendstore "cs_chat_app_server/modules/friend/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListFriend(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		findFriendBiz := friendbiz.NewFindFriendBiz(friendStore)
		friends, err := findFriendBiz.FindFriend(context.Request.Context(), requester.GetId())
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": friends})
	}
}
