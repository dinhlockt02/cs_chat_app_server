package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	usergin "cs_chat_app_server/modules/user/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitUserRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	user := g.Group("/user", authmiddleware.Authentication(appCtx))
	{
		user.PUT("/self", usergin.UpdateSelf(appCtx))
	}
}