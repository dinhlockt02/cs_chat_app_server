package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	authredis "cs_chat_app_server/modules/auth/redis"
	authstore "cs_chat_app_server/modules/auth/store"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func VerifyEmail(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		code, ok := context.GetQuery("code")
		if !ok {
			log.Error().Err(common.ErrInvalidRequest(errors.New("invalid verify url"))).Send()
			context.Redirect(http.StatusFound, "/verify/failure")
			return
		}
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err := authbiz.NewVerifyEmail(authStore, authredis.NewRedisStore(
			appCtx.RedisClient(),
		)).Verify(context.Request.Context(), code)
		if err != nil {
			log.Error().Err(err).Send()
			context.Redirect(http.StatusFound, "/verify/failure")
			return
		}
		context.Redirect(http.StatusFound, "/verify/success")
	}
}
