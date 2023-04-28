package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	authgin "cs_chat_app_server/modules/auth/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	auth := g.Group("/auth")
	{
		auth.POST("/register", authgin.Register(appCtx))
		auth.POST("/login", authgin.Login(appCtx))
		auth.POST("/update-password", authmiddleware.Authentication(appCtx), authgin.UpdatePassword(appCtx))

	}
}
