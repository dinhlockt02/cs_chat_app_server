package middleware

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"cs_chat_app_server/components/socket"
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
						log.Error().Err(appErr)
					}
					return
				}

				appErr := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				if gin.Mode() == gin.DebugMode {
					panic(err)
				} else {
					log.Error().Err(err.(error))
				}
				return
			}
		}()
		c.Next()
	}
}

func RecoverSocket(c *socket.Context) {
	if err := recover(); err != nil {
		if err == context.Canceled {
			return
		}
		if appErr, ok := err.(*common.AppError); ok {
			log.Error().Stack().Err(appErr).Msg("")
			c.Response(appErr)
			return
		}

		appErr := common.ErrInternal(err.(error))
		log.Error().Stack().Err(appErr).Msg("")
		c.Response(appErr)
		return
	}
}
