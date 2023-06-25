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

func UpdateGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		groupId := c.Param("groupId")

		if _, err := common.ToObjectId(groupId); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data *groupmdl.UpdateGroup

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		updateGroupBiz := groupbiz.NewUpdateGroupBiz(groupRepo, appCtx.PubSub())

		if err := updateGroupBiz.Update(c.Request.Context(), requester.GetId(), groupId, data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
