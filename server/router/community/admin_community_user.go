package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AdminCommunityUserRouter struct {
}

// InitAdminCommunityUserRouter 初始化 CommunityUser 路由信息
func (s *AdminCommunityUserRouter) InitAdminCommunityUserRouter(Router *gin.RouterGroup) {
	communityUserRouter := Router.Group("communityUser").Use(middleware.OperationRecord())
	communityUserRouterWithoutRecord := Router.Group("communityUser")
	var communityUserApi = v1.ApiGroupApp.CommunityApiGroup.CommunityUserApi
	{
		communityUserRouter.POST("createCommunityUser", communityUserApi.CreateCommunityUser)             // 新建CommunityUser
		communityUserRouter.DELETE("deleteCommunityUser", communityUserApi.DeleteCommunityUser)           // 删除CommunityUser
		communityUserRouter.DELETE("deleteCommunityUserByIds", communityUserApi.DeleteCommunityUserByIds) // 批量删除CommunityUser
		communityUserRouter.PUT("updateCommunityUser", communityUserApi.UpdateCommunityUser)              // 更新CommunityUser
	}
	{
		communityUserRouterWithoutRecord.GET("findCommunityUser", communityUserApi.FindCommunityUser)       // 根据ID获取CommunityUser
		communityUserRouterWithoutRecord.GET("getCommunityUserList", communityUserApi.GetCommunityUserList) // 获取CommunityUser列表
	}
}
