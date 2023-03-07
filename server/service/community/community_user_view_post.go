package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
)

type UserViewPostService struct {
}

// Add 添加点赞
func (userViewPostService *UserViewPostService) Add(model community.UserViewPost) (ok bool, err error) {
	var data community.UserViewPost
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("post_id = ?", model.PostId).First(&data).Error

	if err != nil {
		return false, errors.New("查询失败")
	}

	err = global.GVA_DB.Create(&model).Error
	return true, err
}
