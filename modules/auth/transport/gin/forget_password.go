package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	authmodel "cs_chat_app_server/modules/auth/model"
	authredis "cs_chat_app_server/modules/auth/redis"
	authstore "cs_chat_app_server/modules/auth/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ForgetPassword(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		type Body struct {
			Email string `json:"email"`
		}
		body := &Body{}

		err := c.ShouldBind(body)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if m := common.EmailRegexp.Match([]byte(body.Email)); !m {
			panic(common.ErrInvalidRequest(authmodel.ErrInvalidEmail))

		}

		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		err = authbiz.
			NewForgetPasswordBiz(
				appCtx.Mailer(),
				authStore,
				authredis.NewRedisStore(
					appCtx.RedisClient(),
				),
			).
			Execute(c.Request.Context(), body.Email)
		if err != nil {
			panic(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
