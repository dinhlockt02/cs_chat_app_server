package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewAccessToken(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		type Body struct {
			RefreshToken string `json:"refresh_token"`
		}

		var body Body

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		biz := authbiz.NewAccessTokenBiz(appCtx.TokenProvider())
		result, err := biz.New(context.Request.Context(), body.RefreshToken)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
