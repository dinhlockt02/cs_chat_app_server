package v1route

import (
	"cs_chat_app_server/components/appcontext"
	routeinternal "cs_chat_app_server/route/v1/internal"
	"github.com/gin-gonic/gin"
)

func InitRoute(e *gin.Engine, appCtx appcontext.AppContext) {
	v1 := e.Group("/v1")
	{
		routeinternal.InitAuthRoute(v1, appCtx)
		routeinternal.InitUserRoute(v1, appCtx)
		routeinternal.InitFriendRoute(v1, appCtx)
		routeinternal.InitSocketRoute(v1, appCtx)
		routeinternal.InitGroupRoute(v1, appCtx)
		routeinternal.InitDeviceRoute(v1, appCtx)
	}
}
