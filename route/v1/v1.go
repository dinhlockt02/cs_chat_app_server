package v1route

import (
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/route/v1/internal"
	"github.com/gin-gonic/gin"
)

func InitRoute(e *gin.Engine, appCtx appcontext.AppContext) {
	v1 := e.Group("/v1")
	{
		internal.InitAuthRoute(v1, appCtx)
	}
}