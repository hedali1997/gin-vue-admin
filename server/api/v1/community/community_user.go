package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	communityRes "github.com/flipped-aurora/gin-vue-admin/server/model/community/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"time"
)

type CommunityUserApi struct {
}

var communityUserService = service.ServiceGroupApp.CommunityServiceGroup.CommunityUserService

// PhoneLogin
// @Tags     CommunityBase
// @Summary  社区用户手机号密码登录
// @Produce   application/json
// @Param    data  body      communityReq.PhoneLogin                                             true  "手机号, 密码, 验证码"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /communityBase/login [post]
func (b *CommunityBaseApi) PhoneLogin(c *gin.Context) {
	var l communityReq.PhoneLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.PhoneLoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 判断验证码是否开启
	openCaptcha := global.GVA_CONFIG.Captcha.OpenCaptcha               // 是否开启防爆次数
	openCaptchaTimeOut := global.GVA_CONFIG.Captcha.OpenCaptchaTimeOut // 缓存超时时间
	v, ok := global.BlackCache.Get(key)
	if !ok {
		global.BlackCache.Set(key, 1, time.Second*time.Duration(openCaptchaTimeOut))
	}

	var oc bool = openCaptcha == 0 || openCaptcha < interfaceToInt(v)

	if !oc || store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &community.CommunityUser{Phone: l.Phone, Password: l.Password}
		user, err := userService.Login(u)
		if err != nil {
			global.GVA_LOG.Error("登陆失败! 手机号不存在或者密码错误!", zap.Error(err))
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("手机号不存在或者密码错误", c)
			return
		}
		if user.Status == 2 {
			global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
			// 验证码次数+1
			global.BlackCache.Increment(key, 1)
			response.FailWithMessage("用户被禁止登录", c)
			return
		}

		// 前端接口仅返回必要的字段
		var tempUser *community.ApiCommunityUser

		err = utils.Copy(&tempUser, user)
		if err != nil {
			global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
			response.FailWithMessage("获取用户信息失败!", c)
			return
		}

		b.TokenNext(c, *tempUser)
		return
	}
	// 验证码次数+1
	global.BlackCache.Increment(key, 1)
	response.FailWithMessage("验证码错误", c)
}

// PhoneCodeLogin
// @Tags     CommunityBase
// @Summary  社区用户短信验证登录
// @Produce   application/json
// @Param    data  body      communityReq.PhoneCodeLogin                                    true  "手机号, 验证码"
// @Success  200   {object}  response.Response{data=communityRes.LoginResponse,msg=string}  "返回包括用户信息,token,过期时间"
// @Router   /communityBase/codeLogin [post]
func (b *CommunityBaseApi) PhoneCodeLogin(c *gin.Context) {
	var l communityReq.PhoneCodeLogin
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(l, utils.PhoneCodeLoginVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	res := verifyCode(c, l.Phone, l.Code, "PhoneCodeLogin")
	if res.Code != 0 {
		response.FailWithMessage(res.Msg, c)
		return
	}

	u := &community.CommunityUser{Phone: l.Phone}
	user, err := userService.Login(u)
	if err != nil {
		global.GVA_LOG.Error("登陆失败! 手机号不存在或者密码错误!", zap.Error(err))
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("手机号不存在或者密码错误", c)
		return
	}
	if user.Status == 2 {
		global.GVA_LOG.Error("登陆失败! 用户被禁止登录!")
		// 验证码次数+1
		global.BlackCache.Increment(key, 1)
		response.FailWithMessage("用户被禁止登录", c)
		return
	}

	// 前端接口仅返回必要的字段
	var tempUser *community.ApiCommunityUser

	err = utils.Copy(&tempUser, user)
	if err != nil {
		global.GVA_LOG.Error("获取用户信息失败!", zap.Error(err))
		response.FailWithMessage("获取用户信息失败!", c)
		return
	}

	b.TokenNext(c, *tempUser)
	return
}

// TokenNext 登录以后签发jwt
func (b *CommunityBaseApi) TokenNext(c *gin.Context, user community.ApiCommunityUser) {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateCommunityClaims(systemReq.CommunityBaseClaims{
		UUID:     user.UUID,
		ID:       user.ID,
		NickName: user.Nickname,
		Phone:    user.Phone,
		Status:   user.Status,
	})
	token, err := j.CreateCommunityToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if jwtStr, err := jwtService.GetRedisJWT(user.Phone); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Phone); err != nil {
			global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
		response.FailWithMessage("设置登录状态失败", c)
	} else {
		var blackJWT system.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Phone); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.OkWithDetailed(communityRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

// Register
// @Tags     communityBase
// @Summary  社区用户注册账号
// @Produce   application/json
// @Param    data  body      communityReq.Register                                        true  "用户名, 昵称, 密码, 角色ID"
// @Success  200   {object}  response.Response{data=communityRes.RegisterResponse,msg=string}  "用户注册账号,返回包括用户信息"
// @Router   /communityBase/register [post]
func (b *CommunityBaseApi) Register(c *gin.Context) {
	var r communityReq.Register
	err := c.ShouldBindJSON(&r)
	if err != nil {
		global.GVA_LOG.Error("注册参数绑定失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(r, utils.PhoneRegisterVerify)
	if err != nil {
		global.GVA_LOG.Error("注册校验失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if r.Password != r.CheckPassword {
		response.FailWithMessage("密码与确认密码不一致", c)
		return
	}

	// 校验code
	verifyRes := verifyCode(c, r.Phone, r.Code, "Register")
	if verifyRes.Code != 0 {
		response.Result(verifyRes.Code, verifyRes.Data, verifyRes.Msg, c)
		return
	}

	user := &community.CommunityUser{UserName: r.UserName, Phone: r.Phone, Password: r.Password}
	userReturn, err := userService.Register(*user)

	// 前端接口仅返回必要的字段
	var tempUser community.ApiCommunityUser

	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(communityRes.RegisterResponse{User: tempUser}, "注册失败!"+err.Error(), c)
		return
	}

	err = utils.Copy(&tempUser, userReturn)
	if err != nil {
		global.GVA_LOG.Error("注册失败!", zap.Error(err))
		response.FailWithDetailed(communityRes.RegisterResponse{User: tempUser}, "注册失败", c)
		return
	}

	response.OkWithDetailed(communityRes.RegisterResponse{User: tempUser}, "注册成功", c)
}

// ChangePassword
// @Tags      communityBase
// @Summary   社区用户修改密码
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      communityReq.ChangePasswordReq    true  "用户名, 原密码, 新密码"
// @Success   200   {object}  response.Response{msg=string}  "用户修改密码"
// @Router    /communityBase/changePassword [post]
func (b *CommunityBaseApi) ChangePassword(c *gin.Context) {
	var req communityReq.ChangePasswordReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.ChangePasswordVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	uid := utils.GetCommunityUserID(c)
	u := &community.CommunityUser{GVA_MODEL: global.GVA_MODEL{ID: uid}, Password: req.Password}
	_, err = userService.ChangePassword(u, req.NewPassword)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败，原密码与当前账户不符", c)
		return
	}
	response.OkWithMessage("修改成功", c)
}

// CreateCommunityUser 创建CommunityUser
// @Tags CommunityUser
// @Summary 创建CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.CommunityUser true "创建CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /communityUser/createCommunityUser [post]
func (communityUserApi *CommunityUserApi) CreateCommunityUser(c *gin.Context) {
	var communityUser community.CommunityUser
	err := c.ShouldBindJSON(&communityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	communityUser.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserName": {utils.NotEmpty()},
	}
	if err := utils.Verify(communityUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := communityUserService.CreateCommunityUser(communityUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCommunityUser 删除CommunityUser
// @Tags CommunityUser
// @Summary 删除CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.CommunityUser true "删除CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /communityUser/deleteCommunityUser [delete]
func (communityUserApi *CommunityUserApi) DeleteCommunityUser(c *gin.Context) {
	var communityUser community.CommunityUser
	err := c.ShouldBindJSON(&communityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	communityUser.DeletedBy = utils.GetUserID(c)
	if err := communityUserService.DeleteCommunityUser(communityUser); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCommunityUserByIds 批量删除CommunityUser
// @Tags CommunityUser
// @Summary 批量删除CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /communityUser/deleteCommunityUserByIds [delete]
func (communityUserApi *CommunityUserApi) DeleteCommunityUserByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := communityUserService.DeleteCommunityUserByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCommunityUser 更新CommunityUser
// @Tags CommunityUser
// @Summary 更新CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body community.CommunityUser true "更新CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /communityUser/updateCommunityUser [put]
func (communityUserApi *CommunityUserApi) UpdateCommunityUser(c *gin.Context) {
	var communityUser community.CommunityUser
	err := c.ShouldBindJSON(&communityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	communityUser.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"UserName": {utils.NotEmpty()},
	}
	if err := utils.Verify(communityUser, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := communityUserService.UpdateCommunityUser(communityUser); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCommunityUser 用id查询CommunityUser
// @Tags CommunityUser
// @Summary 用id查询CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query community.CommunityUser true "用id查询CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /communityUser/findCommunityUser [get]
func (communityUserApi *CommunityUserApi) FindCommunityUser(c *gin.Context) {
	var communityUser community.CommunityUser
	err := c.ShouldBindQuery(&communityUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recommunityUser, err := communityUserService.GetCommunityUser(communityUser.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recommunityUser": recommunityUser}, c)
	}
}

// GetCommunityUserList 分页获取CommunityUser列表
// @Tags CommunityUser
// @Summary 分页获取CommunityUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query communityReq.CommunityUserSearch true "分页获取CommunityUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /communityUser/getCommunityUserList [get]
func (communityUserApi *CommunityUserApi) GetCommunityUserList(c *gin.Context) {
	var pageInfo communityReq.CommunityUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := communityUserService.GetCommunityUserInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
