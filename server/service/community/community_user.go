package community

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	communityReq "github.com/flipped-aurora/gin-vue-admin/server/model/community/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type CommunityUserService struct {
}

func (communityUserService *CommunityUserService) Register(u community.CommunityUser) (userInter community.CommunityUser, err error) {
	var user community.CommunityUser
	if !errors.Is(global.GVA_DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return userInter, errors.New("用户名已被使用")
	}

	if !errors.Is(global.GVA_DB.Where("phone = ?", u.Phone).First(&user).Error, gorm.ErrRecordNotFound) { // 判断手机号是否注册
		return userInter, errors.New("手机号已注册")
	}

	u.Sex = 3
	u.Status = 1
	u.Education = 1
	// 否则 附加uuid 密码hash加密 注册
	u.Password = utils.BcryptHash(u.Password)
	u.UUID = uuid.NewV4()
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: Login
//@description: 用户登录
//@param: u *community.CommunityUser
//@return: err error, userInter *community.CommunityUser

func (communityUserService *CommunityUserService) Login(u *community.CommunityUser) (userInter *community.CommunityUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user community.CommunityUser
	err = global.GVA_DB.Where("phone = ?", u.Phone).First(&user).Error
	if err == nil {
		if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
			return nil, errors.New("密码错误")
		}
		//MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

func (communityUserService *CommunityUserService) CodeLogin(u *community.CommunityUser) (userInter *community.CommunityUser, err error) {
	if nil == global.GVA_DB {
		return nil, fmt.Errorf("db not init")
	}

	var user community.CommunityUser
	err = global.GVA_DB.Where("phone = ?", u.Phone).First(&user).Error
	return &user, err
}

//@author: [hedali](https://github.com/hedali1997)
//@function: ChangePassword
//@description: 修改社区用户密码
//@param: u *community.CommunityUser, newPassword string
//@return: userInter *community.CommunityUser,err error

func (communityUserService *CommunityUserService) ChangePassword(u *community.CommunityUser, newPassword string) (userInter *community.CommunityUser, err error) {
	var user community.CommunityUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.Password, user.Password); !ok {
		return nil, errors.New("原密码错误")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

func (communityUserService *CommunityUserService) ChangeInfo(u *community.CommunityUser) (userInter *community.CommunityUser, err error) {
	var user community.CommunityUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}

	user.Avatar = u.Avatar
	user.Nickname = u.Nickname
	user.Sex = u.Sex
	user.School = u.School
	user.Education = u.Education
	user.Major = u.Major

	if u.Birthday != "" {
		user.Birthday = u.Birthday
	}

	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

func (communityUserService *CommunityUserService) RecoverPassword(u *community.CommunityUser, newPassword string) (userInter *community.CommunityUser, err error) {
	var user community.CommunityUser
	if err = global.GVA_DB.Where("phone = ?", u.Phone).First(&user).Error; err != nil {
		return nil, err
	}
	if user.ID == 0 {
		return nil, errors.New("用户不存在")
	}
	user.Password = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err

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
