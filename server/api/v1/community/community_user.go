package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CommunityUserApi struct {
}

var communityUserService = service.ServiceGroupApp.CommunityServiceGroup.CommunityUserService

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
