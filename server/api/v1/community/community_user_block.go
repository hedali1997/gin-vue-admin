package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserBlockApi struct {
}

var userBlockService = service.ServiceGroupApp.CommunityServiceGroup.UserBlockService

// Block
// @Tags      communityUserBlock
// @Summary   拉黑用户
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      communityReq.EditUserBlockReq    true  "用户id，被拉黑的用户id, 操作？1拉黑2取消拉黑"
// @Success   200   {object}  response.Response{msg=string}  "是否成功"
// @Router    /communityUser/Block [post]
func (b *UserBlockApi) Block(c *gin.Context) {
	var req communityReq.EditUserBlockReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.UserBlockVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := community.UserBlock{UserId: req.UserId, BlockUserId: req.BlockUserId}

	if req.IsBlock == 2 {
		_, err = userBlockService.Del(data)
	} else {
		_, err = userBlockService.Add(data)
	}

	if err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败！"+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
