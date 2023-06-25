package searchgin

import (
	"context"
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendbiz "cs_chat_app_server/modules/friend/biz"
	friendmodel "cs_chat_app_server/modules/friend/model"
	friendstore "cs_chat_app_server/modules/friend/store"
	groupbiz "cs_chat_app_server/modules/group/biz"
	groupmdl "cs_chat_app_server/modules/group/model"
	grouprepo "cs_chat_app_server/modules/group/repository"
	groupstore "cs_chat_app_server/modules/group/store"
	requeststore "cs_chat_app_server/modules/request/store"

	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

func Search(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		searchTerm := c.Query("term")

		u, _ := c.Get(common.CurrentUser)
		requester := u.(common.Requester)

		wg := sync.WaitGroup{}
		wg.Add(2)
		rs := map[string]interface{}{}
		go func() {
			defer common.Recovery()
			defer wg.Done()
			friends, err := searchFriend(c.Request.Context(), appCtx, requester, searchTerm)
			if err != nil {
				panic(err)
			}
			rs["friends"] = friends
		}()
		go func() {
			defer common.Recovery()
			defer wg.Done()
			groups, err := searchGroup(c.Request.Context(), appCtx, requester, searchTerm)
			if err != nil {
				panic(err)
			}
			rs["groups"] = groups
		}()

		wg.Wait()

		c.JSON(http.StatusOK, gin.H{"data": rs})
	}
}

func searchFriend(ctx context.Context, appCtx appcontext.AppContext, requester common.Requester, searchTerm string) ([]friendmodel.FriendUser, error) {
	friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupRepo := grouprepo.NewGroupRepository(groupStore, requestStore)
	findFriendBiz := friendbiz.NewFindFriendBiz(friendStore, groupRepo)

	filter := map[string]interface{}{}
	err := common.AddIdFilter(map[string]interface{}{}, requester.GetId())
	if err != nil {
		panic(err)
	}
	friends, err := findFriendBiz.FindFriend(ctx, requester.GetId(), common.GetAndFilter(filter, common.GetTextSearch(searchTerm, false, false)))
	if err != nil {
		panic(err)
	}
	return friends, nil
}

func searchGroup(ctx context.Context, appCtx appcontext.AppContext, requester common.Requester, searchTerm string) ([]groupmdl.Group, error) {
	groupStore := groupstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	requestStore := requeststore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
	groupRepo := grouprepo.NewGroupRepository(
		groupStore,
		requestStore,
	)
	listGroupBiz := groupbiz.NewListGroupBiz(groupRepo)
	groups, err := listGroupBiz.List(ctx, requester.GetId(), common.GetTextSearch(searchTerm, false, true))
	if err != nil {
		panic(err)
	}
	return groups, nil
}
