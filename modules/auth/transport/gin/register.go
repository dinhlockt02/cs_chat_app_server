package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	authmodel "cs_chat_app_server/modules/auth/model"
	authstore "cs_chat_app_server/modules/auth/store"
	devicemodel "cs_chat_app_server/modules/device/model"
	devicestore "cs_chat_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		type Body struct {
			Data   authmodel.RegisterUser `json:"data"`
			Device devicemodel.Device     `json:"device"`
		}

		var body = Body{
			Data:   authmodel.RegisterUser{},
			Device: devicemodel.Device{},
		}

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewRegisterBiz(appCtx.TokenProvider(), appCtx.Hasher(), deviceStore, authStore)
		result, err := biz.Register(context.Request.Context(), &body.Data, &body.Device)
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
