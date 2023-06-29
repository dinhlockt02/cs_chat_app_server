package notigin

import (
	notirepo "cs_chat_app_server/components/notification/repository"
	notistore "cs_chat_app_server/components/notification/store"
	notibiz "cs_chat_app_server/modules/notification/biz"
	"net/http"

	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	"github.com/gin-gonic/gin"
)

func DeleteNotificationById(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		id := context.Param("notificationId")

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		store := notistore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := notirepo.NewNotificationRepository(store)
		biz := notibiz.NewDeleteByIdNotificationBiz(repo)
		err := biz.DeleteById(context.Request.Context(), requesterId, id)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": true})
	}
}
