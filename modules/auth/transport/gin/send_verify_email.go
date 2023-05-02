package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	authredis "cs_chat_app_server/modules/auth/redis"
	authstore "cs_chat_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendVerifyEmail(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err := authbiz.
			NewSendVerifyEmail(
				appCtx.Mailer(),
				authStore,
				authredis.NewRedisStore(
					appCtx.RedisClient(),
				),
			).
			Send(context.Request.Context(), requester.GetId(), false)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
