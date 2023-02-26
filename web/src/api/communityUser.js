import service from '@/utils/request'

// @Tags CommunityUser
// @Summary 创建CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommunityUser true "创建CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /communityUser/createCommunityUser [post]
export const createCommunityUser = (data) => {
  return service({
    url: '/communityUser/createCommunityUser',
    method: 'post',
    data
  })
}

// @Tags CommunityUser
// @Summary 删除CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommunityUser true "删除CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /communityUser/deleteCommunityUser [delete]
export const deleteCommunityUser = (data) => {
  return service({
    url: '/communityUser/deleteCommunityUser',
    method: 'delete',
    data
  })
}

// @Tags CommunityUser
// @Summary 删除CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /communityUser/deleteCommunityUser [delete]
export const deleteCommunityUserByIds = (data) => {
  return service({
    url: '/communityUser/deleteCommunityUserByIds',
    method: 'delete',
    data
  })
}

// @Tags CommunityUser
// @Summary 更新CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CommunityUser true "更新CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /communityUser/updateCommunityUser [put]
export const updateCommunityUser = (data) => {
  return service({
    url: '/communityUser/updateCommunityUser',
    method: 'put',
    data
  })
}

// @Tags CommunityUser
// @Summary 用id查询CommunityUser
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CommunityUser true "用id查询CommunityUser"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /communityUser/findCommunityUser [get]
export const findCommunityUser = (params) => {
  return service({
    url: '/communityUser/findCommunityUser',
    method: 'get',
    params
  })
}

// @Tags CommunityUser
// @Summary 分页获取CommunityUser列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CommunityUser列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /communityUser/getCommunityUserList [get]
export const getCommunityUserList = (params) => {
  return service({
    url: '/communityUser/getCommunityUserList',
    method: 'get',
    params
  })
}
