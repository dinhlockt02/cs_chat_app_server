package gchatgin

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	groupstore "cs_chat_app_server/modules/group/store"
	gchatskt "cs_chat_app_server/modules/group_chat/transport/socket"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gobwas/ws"
	"github.com/rs/zerolog/log"
)

func GroupWebsocketChatHandler(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		u, _ := c.Get(common.CurrentUser)
		requester, _ := u.(common.Requester)
		ctx = context.WithValue(ctx, common.CurrentUser, u)

		groupId := c.Param("groupId")
		groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		filter, err := common.GetIdFilter(groupId)

		if err != nil {
			log.Debug().Err(err).Str("package", "gchatgin.GroupWebsocketChatHandler").
				Msg("can not get id filter from group id: " + groupId)
			panic(err)
		}
		filter = common.GetAndFilter(
			filter,
			groupstore.GetMemberIdInGroupMembersFilter(requester.GetId()))

		group, err := groupStore.FindGroup(c.Request.Context(), filter)
		if err != nil {
			log.Error().Err(err).Str("package", "gchatgin.GroupWebsocketChatHandler").Send()
			panic(err)
		}

		if group == nil {
			err := errors.New("group not found")
			log.Debug().Err(err).Str("package", "gchatgin.GroupWebsocketChatHandler").Send()
			panic(common.ErrInvalidRequest(err))
		}
		ctx = context.WithValue(ctx, common.CurrentGroupId, groupId)

		conn, _, _, err := ws.UpgradeHTTP(c.Request, c.Writer)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		err = appCtx.Socket().AddConn(requester.GetId(), conn)
		if err != nil {
			panic(common.ErrInternal(err))
		}

		appCtx.Socket().Receive(conn, ctx, gchatskt.SendMessageHandler(appCtx))
	}
}
