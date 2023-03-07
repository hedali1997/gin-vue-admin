package community

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	CommunityUserApi
	CommunityBaseApi
	UserBlockApi
	UserCollectPostApi
	UserFenApi
}

var (
	userService = service.ServiceGroupApp.CommunityServiceGroup.CommunityUserService
	jwtService  = service.ServiceGroupApp.SystemServiceGroup.JwtService
)
