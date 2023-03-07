package community

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CommunityUserRouter struct {
}

func (s *CommunityBaseRouter) InitCommunityUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("communityUser")
	userApi := v1.ApiGroupApp.CommunityApiGroup.CommunityUserApi
	userBlockApi := v1.ApiGroupApp.CommunityApiGroup.UserBlockApi
	collectApi := v1.ApiGroupApp.CommunityApiGroup.UserCollectPostApi
	fenApi := v1.ApiGroupApp.CommunityApiGroup.UserFenApi
	{
		baseRouter.POST("editUser", userApi.EditUser)  // 修改个人信息
		baseRouter.POST("Block", userBlockApi.Block)   // 拉黑
		baseRouter.POST("Collect", collectApi.Collect) // 收藏
		baseRouter.POST("Fen", fenApi.Fen)             // 收藏
	}
	return baseRouter
}
