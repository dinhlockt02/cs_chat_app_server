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
	"github.com/rs/zerolog/log"
	"net/http"
)

func ListFriend(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(groupStore, requestStore)
		findFriendBiz := friendbiz.NewFindFriendBiz(friendStore, groupRepo)
		friends, err := findFriendBiz.FindFriend(context.Request.Context(), requester.GetId(), map[string]interface{}{})
		if err != nil {
			log.Error().Err(err).
				Str("package", "friendgin.ListFriend").
				Msgf("error while call findFriendBiz.FindFriend")
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": friends})
	}
}
