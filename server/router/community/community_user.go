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
	{
		baseRouter.POST("editUser", userApi.EditUser) // 修改个人信息
	}
	return baseRouter
}
