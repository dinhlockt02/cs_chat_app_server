package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	searchgin "cs_chat_app_server/modules/search/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitSearchRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	g.GET("/search", authmiddleware.Authentication(appCtx), searchgin.Search(appCtx))
}
