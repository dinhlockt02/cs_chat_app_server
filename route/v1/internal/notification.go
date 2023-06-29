package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	notigin "cs_chat_app_server/modules/notification/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitNotificationRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {
	notification := g.Group("/notification", authmiddleware.Authentication(appCtx))
	{
		notification.GET("", notigin.ListNotification(appCtx))
		notification.DELETE("", notigin.DeleteAllNotifications(appCtx))
		notification.DELETE("/:notificationId", notigin.DeleteNotificationById(appCtx))
	}
}
