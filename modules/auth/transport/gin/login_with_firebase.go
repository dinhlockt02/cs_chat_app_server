package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	authstore "cs_chat_app_server/modules/auth/store"
	devicemodel "cs_chat_app_server/modules/device/model"
	devicestore "cs_chat_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginWithFirebase(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		type Body struct {
			IdToken string             `json:"id_token"`
			Device  devicemodel.Device `json:"device"`
		}

		var body = Body{
			IdToken: "",
			Device:  devicemodel.Device{},
		}

		if err := context.ShouldBind(&body); err != nil {
			panic(common.ErrInvalidRequest(err))
			return
		}

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		authStore := authstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewLoginWithFirebaseBiz(appCtx.TokenProvider(), deviceStore, authStore, appCtx.FirebaseApp())
		result, err := biz.LoginWithFirebase(context.Request.Context(), body.IdToken, &body.Device)

		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": result})
	}
}
