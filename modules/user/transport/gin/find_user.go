package usergin

import (
	"cs_chat_app_server/common"
	"cs_chat_app_server/components/appcontext"
	friendrepo "cs_chat_app_server/modules/friend/repository"
	friendstore "cs_chat_app_server/modules/friend/store"
	userbiz "cs_chat_app_server/modules/user/biz"
	userrepo "cs_chat_app_server/modules/user/repository"
	userstore "cs_chat_app_server/modules/user/store"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func FindUser(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(context *gin.Context) {
		u, _ := context.Get(common.CurrentUser)
		requester := u.(common.Requester)

		type filterBody struct {
			Email string `form:"email,omitempty"`
		}

		var filter filterBody

		err := context.ShouldBind(&filter)

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		userStore := userstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))
		friendStore := friendstore.NewMongoStore(appCtx.MongoClient().Database(common.AppDatabase))

		friendRepo := friendrepo.NewFindUserRepository(friendStore)
		findUserRepo := userrepo.NewFindUserRepo(userStore, friendRepo)

		findUserBiz := userbiz.NewFindUserBiz(findUserRepo)

		id, err := primitive.ObjectIDFromHex(requester.GetId())
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		user, err := findUserBiz.FindUser(context.Request.Context(), id.Hex(), map[string]interface{}{
			"email": filter.Email,
		})
		if err != nil {
			panic(err)
		}

		context.JSON(http.StatusOK, gin.H{"data": user})
	}
}
