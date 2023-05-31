package groupgin

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupbiz "cs_chat_app_server/modules/group/biz"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	requeststore "cs_chat_app_server/modules/request/store"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"sync"
)

func CreateGroup(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data *groupmdl.Group

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		invitedMembers := data.Members

		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		groupRepo := grouprepo.NewGroupRepository(
			groupStore,
			requestStore,
		)
		createGroupBiz := groupbiz.NewCreateGroupBiz(groupRepo)

		if err := createGroupBiz.Create(c.Request.Context(), requester.GetId(), data); err != nil {
			panic(err)
		}

		go func() {
			defer common.Recovery()
			wg := sync.WaitGroup{}
			sendGroupRequestBiz := groupbiz.NewSendGroupRequestBiz(groupRepo)
			for _, member := range invitedMembers {
				if member != requester.GetId() {
					go func() {
						wg.Add(1)
						defer wg.Done()
						defer common.Recovery()
						err := sendGroupRequestBiz.SendRequest(context.Background(), requester.GetId(), member, data)
						if err != nil {
							log.Error().Msgf("%v\n", err)
						}
					}()
				}
			}
			wg.Wait()
		}()

		c.JSON(http.StatusCreated, gin.H{"data": data.Id})
	}
}
