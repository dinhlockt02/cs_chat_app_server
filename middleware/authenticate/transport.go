package authmiddleware

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Authentication(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		authorizationHeader := strings.Split(c.GetHeader("Authorization"), " ")

		if len(authorizationHeader) != 2 || authorizationHeader[0] != "Bearer" {
			var unauthorizedError = common.NewFullErrorResponse(http.StatusUnauthorized,
				nil,
				"unauthorized",
				"Invalid header",
				"UnauthorizedError")

			panic(unauthorizedError)
		}

		store := NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := NewAuthMiddlewareBiz(store, appCtx.TokenProvider())
		device, err := biz.Authenticate(c.Request.Context(), authorizationHeader[1])
		if err != nil {
			panic(err)
		}
		c.Set(common.CurrentUser, device)
		c.Next()
	}
}
