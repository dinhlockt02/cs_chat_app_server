package internal

import (
	"cs_chat_app_server/components/appcontext"
	authgin "cs_chat_app_server/modules/auth/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitAuthRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	auth := g.Group("/auth")
	{
		auth.POST("/register", authgin.Register(appCtx))
	}
}
