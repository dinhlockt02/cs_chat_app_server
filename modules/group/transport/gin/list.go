package groupgin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupbiz "cs_chat_app_server/modules/group/biz"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	requeststore "cs_chat_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var query groupmdl.ListGroupQuery

		err := c.ShouldBind(&query)
		if err != nil {
			return
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		listGroupBiz := groupbiz.NewListGroupBiz(groupRepo)
		groups, err := listGroupBiz.List(c.Request.Context(), requester.GetId(), query.ToMap())
		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": groups})
	}
}
