package devicegin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	devicebiz "cs_chat_app_server/modules/device/biz"
	devicestore "cs_chat_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetDevices(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		deviceStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		devices, err := devicebiz.NewGetDevicesBiz(deviceStore).Get(context.Request.Context(), devicestore.GetUserIdFilter(requester.GetId()))
		if err != nil {
			panic(err)
		}
		context.JSON(http.StatusOK, gin.H{"data": devices})
	}
}
