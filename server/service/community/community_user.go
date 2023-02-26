package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"gorm.io/gorm"
)

type CommunityUserService struct {
}

// CreateCommunityUser 创建CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) CreateCommunityUser(communityUser community.CommunityUser) (err error) {
	err = global.GVA_DB.Create(&communityUser).Error
	return err
}

// DeleteCommunityUser 删除CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) DeleteCommunityUser(communityUser community.CommunityUser) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&community.CommunityUser{}).Where("id = ?", communityUser.ID).Update("deleted_by", communityUser.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&communityUser).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCommunityUserByIds 批量删除CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) DeleteCommunityUserByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&community.CommunityUser{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&community.CommunityUser{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCommunityUser 更新CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) UpdateCommunityUser(communityUser community.CommunityUser) (err error) {
	err = global.GVA_DB.Save(&communityUser).Error
	return err
}

// GetCommunityUser 根据id获取CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) GetCommunityUser(id uint) (communityUser community.CommunityUser, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&communityUser).Error
	return
}

// GetCommunityUserInfoList 分页获取CommunityUser记录
// Author [piexlmax](https://github.com/piexlmax)
func (communityUserService *CommunityUserService) GetCommunityUserInfoList(info communityReq.CommunityUserSearch) (list []community.CommunityUser, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&community.CommunityUser{})
	var communityUsers []community.CommunityUser
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.UserName != "" {
		db = db.Where("user_name LIKE ?", "%"+info.UserName+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Email != "" {
		db = db.Where("email LIKE ?", "%"+info.Email+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&communityUsers).Error
	return communityUsers, total, err
}
