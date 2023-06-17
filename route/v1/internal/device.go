package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	devicegin "cs_chat_app_server/modules/device/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitDeviceRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	user := g.Group("/device", authmiddleware.Authentication(appCtx))
	{
		user.PUT("", devicegin.UpdateDevice(appCtx))
		user.GET("", devicegin.GetDevices(appCtx))
		user.DELETE("/:deviceId", devicegin.DeleteDevice(appCtx))

	}
}
