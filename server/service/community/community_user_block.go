package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	_ "gorm.io/gorm"
)

type UserBlockService struct {
}

func (userBlockService *UserBlockService) Add(u community.UserBlock) (ok bool, err error) {
	var data community.UserBlock
	err = global.GVA_DB.Where("user_id = ?", u.UserId).Where("block_user_id = ?", u.BlockUserId).First(&data).Error

	if err != nil { // 判断用户名是否注册
		return false, errors.New("用户名已被使用")
	}

	err = global.GVA_DB.Create(&u).Error
	return true, err
}

func (userBlockService *UserBlockService) Del(u community.UserBlock) (ok bool, err error) {
	var data community.UserBlock
	err = global.GVA_DB.Where("user_id = ?", u.UserId).Where("block_user_id = ?", u.BlockUserId).First(&data).Error
	if err != nil { // 判断用户名是否注册
		return false, errors.New("用户名已被使用")
	}

	global.GVA_DB.Where("user_id = ?", u.UserId).Where("block_user_id = ?", u.BlockUserId).Delete(&u)

	return true, err
}
