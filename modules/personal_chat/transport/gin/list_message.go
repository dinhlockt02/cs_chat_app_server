package pchatgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	pchatbiz "cs_chat_app_server/modules/personal_chat/biz"
	pchatmdl "cs_chat_app_server/modules/personal_chat/model"
	pchatstore "cs_chat_app_server/modules/personal_chat/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ListMessage(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		var paging pchatmdl.Paging

		err := context.ShouldBind(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userId := context.Param("id")

		if !primitive.IsValidObjectID(userId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}
		if !primitive.IsValidObjectID(requesterId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		filter := map[string]interface{}{
			"$or": []map[string]interface{}{
				{
					"$and": []map[string]interface{}{
						{"sender": requesterId},
						{"receiver": userId},
					},
				},
				{
					"$and": []map[string]interface{}{
						{"sender": userId},
						{"receiver": requesterId},
					},
				},
			},
		}

		store := pchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		biz := pchatbiz.NewListMessageBiz(store)
		list, err := biz.List(context.Request.Context(), filter, paging)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": list})
	}
}
