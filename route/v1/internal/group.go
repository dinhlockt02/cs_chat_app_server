package routeinternal

import (
	"cs_chat_app_server/components/appcontext"
	authmiddleware "cs_chat_app_server/middleware/authenticate"
	groupgin "cs_chat_app_server/modules/group/transport/gin"
	"github.com/gin-gonic/gin"
)

func InitGroupRoute(g *gin.RouterGroup, appCtx appcontext.AppContext) {

	group := g.Group("/group", authmiddleware.Authentication(appCtx))
	{
		groupRequest := group.Group("/request")
		{
			groupRequest.GET("/sent", groupgin.GetSentRequest(appCtx))
			groupRequest.GET("/received", groupgin.GetReceiveRequest(appCtx))
			groupRequest.POST("/:groupId/accept", groupgin.AcceptRequest(appCtx))
			groupRequest.DELETE("/:groupId/reject", groupgin.RejectRequest(appCtx))

			groupRequest.POST("/:groupId/:friendId", groupgin.SendGroupRequest(appCtx))
			groupRequest.DELETE("/:groupId/:friendId", groupgin.RecallRequest(appCtx))
		}
		group.POST("", groupgin.CreateGroup(appCtx))
		group.GET("", groupgin.ListGroup(appCtx))
	}
}
