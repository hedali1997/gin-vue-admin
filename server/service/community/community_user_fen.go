package community

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
)

type UserFenService struct {
}

// Add 添加关注
func (userFenService *UserFenService) Add(model community.UserFen) (ok bool, err error) {
	var data community.UserFen
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("fen_user_id = ?", model.FenUserId).First(&data).Error

	if err != nil {
		return false, errors.New("查询失败")
	}

	err = global.GVA_DB.Create(&model).Error
	return true, err
}

// Del 移除关注
func (userFenService *UserFenService) Del(model community.UserFen) (ok bool, err error) {
	var data community.UserFen
	err = global.GVA_DB.Where("user_id = ?", model.UserId).Where("fen_user_id = ?", model.FenUserId).First(&data).Error
	if err != nil {
		return false, errors.New("查询失败")
	}

	global.GVA_DB.Where("user_id = ?", model.UserId).Where("fen_user_id = ?", model.FenUserId).Delete(&model)

	return true, err
}
