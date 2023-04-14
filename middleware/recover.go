package middleware

import (
	"context"
	"github.com/dinhlockt02/cs_video_call_app_server/common"
	"github.com/dinhlockt02/cs_video_call_app_server/components/appcontext"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
)

func Recover(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				if err == context.Canceled {
					return
				}
				c.Header("Content-Type", "application/json")
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					if gin.Mode() == gin.DebugMode {
						panic(err)
					} else if appErr.StatusCode >= http.StatusInternalServerError {
						log.Error().Msg(appErr.RootError().Error())
					}
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				if gin.Mode() == gin.DebugMode {
					panic(err)
				} else {
					log.Error().Msg(err.(error).Error())
				}
				return
			}
		}()
		c.Next()
	}
}
