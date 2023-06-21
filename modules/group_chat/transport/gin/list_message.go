package gchatgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	gchatbiz "cs_chat_app_server/modules/group_chat/biz"
	gchatmdl "cs_chat_app_server/modules/group_chat/model"
	gchatrepo "cs_chat_app_server/modules/group_chat/repository"
	gchatstore "cs_chat_app_server/modules/group_chat/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func ListMessage(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {

		var typeFilter map[string]interface{}
		mt := context.Query("type")

		if len(mt) > 0 {
			typeFilter = gchatstore.GetMessageTypeFilter(mt)
		}

		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)
		requesterId := requester.GetId()

		var paging gchatmdl.Paging

		err := context.ShouldBind(&paging)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		groupId := context.Param("groupId")

		if !primitive.IsValidObjectID(groupId) {
			panic(common.ErrInvalidRequest(common.ErrInvalidObjectId))
		}

		filter := map[string]interface{}{
			"group": groupId,
		}

		if typeFilter != nil {
			filter = common.GetAndFilter(filter, typeFilter)
		}

		store := gchatstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		repo := gchatrepo.NewGroupChatRepository(store)
		biz := gchatbiz.NewListMessageBiz(repo)
		list, err := biz.List(context.Request.Context(), requesterId, filter, paging)

		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": list})
	}
}
