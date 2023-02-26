package community

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type CommunityBaseRouter struct{}

func (s *CommunityBaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("communityBase")
	baseApi := v1.ApiGroupApp.CommunityApiGroup.CommunityBaseApi
	{
		baseRouter.POST("login", baseApi.PhoneLogin)
		baseRouter.POST("codeLogin", baseApi.PhoneCodeLogin)
		baseRouter.POST("register", baseApi.Register)
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("sendCode", baseApi.SendCode)
	}
	return baseRouter
}
