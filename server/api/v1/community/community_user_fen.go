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

type UserFenApi struct {
}

var UserFenService = service.ServiceGroupApp.CommunityServiceGroup.UserFenService

// Fen
// @Tags      communityUser
// @Summary   关注/取消关注
// @Security  ApiKeyAuth
// @Produce  application/json
// @Param     data  body      communityReq.EditUserCollectReq    true  "用户id，被拉黑的用户id, 操作？1拉黑2取消拉黑"
// @Success   200   {object}  response.Response{msg=string}  "是否成功"
// @Router    /communityUser/Fen [post]
func (b *UserFenApi) Fen(c *gin.Context) {
	var req communityReq.EditUserFenReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(req, utils.UserFenVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data := community.UserFen{UserId: req.UserId, FenUserId: req.FenUserId}

	if req.IsFen == 2 {
		_, err = UserFenService.Del(data)
	} else {
		_, err = UserFenService.Add(data)
	}

	if err != nil {
		global.GVA_LOG.Error("操作失败!", zap.Error(err))
		response.FailWithMessage("操作失败！"+err.Error(), c)
		return
	}
	response.OkWithMessage("操作成功", c)
}
