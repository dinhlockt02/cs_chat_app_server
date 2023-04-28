package authgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	authbiz "cs_chat_app_server/modules/auth/biz"
	devicestore "cs_chat_app_server/modules/device/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func Logout(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		id, _ := primitive.ObjectIDFromHex(requester.GetDeviceId())

		deviteStore := devicestore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := authbiz.NewLogoutBiz(deviteStore)
		err := biz.Logout(context.Request.Context(), map[string]interface{}{
			"_id": id,
		})
		if err != nil {
			panic(err)
			return
		}
		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
