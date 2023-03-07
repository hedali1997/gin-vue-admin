package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
)

type UserCollectPostService struct {
}

// Add 添加收藏
func (userCollectPostService *UserCollectPostService) Add(model community.UserCollectPost) (ok bool, err error) {
	var data community.UserCollectPost
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).First(&data).Error

	if err != nil {
		return false, errors.New("查询失败")
	}

	err = global.GVA_DB.Create(&model).Error
	return true, err
}

// Del 移除收藏
func (userCollectPostService *UserCollectPostService) Del(model community.UserCollectPost) (ok bool, err error) {
	var data community.UserCollectPost
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).First(&data).Error
	if err != nil {
		return false, errors.New("查询失败")
	}

	global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).Delete(&model)

	return true, err
}
