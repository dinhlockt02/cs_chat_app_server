package usergin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	userbiz "cs_chat_app_server/modules/user/biz"
	usermodel "cs_chat_app_server/modules/user/model"
	userstore "cs_chat_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func UpdateSelf(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		var updateData usermodel.UpdateUser

		err := context.ShouldBind(&updateData)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		updateUserBiz := userbiz.NewUpdateUserBiz(userStore, appCtx.PubSub())

		id, err := primitive.ObjectIDFromHex(requester.GetId())
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err = updateUserBiz.Update(context.Request.Context(), map[string]interface{}{
			"_id": id,
		}, &updateData); err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
