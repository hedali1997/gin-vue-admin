package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
)

type UserStarPostService struct {
}

// Add 添加点赞
func (userStartPostService *UserStarPostService) Add(model community.UserStarPost) (ok bool, err error) {
	var data community.UserStarPost
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).First(&data).Error

	if err != nil {
		return false, errors.New("查询失败")
	}

	err = global.GVA_DB.Create(&model).Error
	return true, err
}

// Del 移除点赞
func (userStartPostService *UserStarPostService) Del(model community.UserStarPost) (ok bool, err error) {
	var data community.UserStarPost
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).First(&data).Error
	if err != nil {
		return false, errors.New("查询失败")
	}

	global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).Delete(&model)

	return true, err
}
