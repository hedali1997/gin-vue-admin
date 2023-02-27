// 自动生成模板CommunityUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/satori/go.uuid"
)

// CommunityUser 结构体
type CommunityUser struct {
	global.GVA_MODEL
	UUID      uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:255;"`
	Nickname  string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:60;"`
	UserName  string    `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:60;"`
	Password  string    `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
	Birthday  string    `json:"birthday" form:"birthday" gorm:"type:date;column:birthday;comment:;"`
	School    string    `json:"school" form:"school" gorm:"column:school;comment:;"`
	Sex       uint8     `json:"sex" form:"sex" gorm:"column:sex;comment:性别：1男2女3未知;"`
	Education uint8     `json:"education" form:"education" gorm:"column:education;comment:;"`
	Status    uint8     `json:"status" form:"status" gorm:"column:status;comment:状态：1正常2禁止登录3禁言;"`
	Major     string    `json:"major" form:"major" gorm:"column:major;comment:;"`
	Avatar    string    `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;size:500;"`
	Phone     string    `json:"phone" form:"phone" gorm:"column:phone;comment:;size:20;"`
	Email     string    `json:"email" form:"email" gorm:"column:email;comment:;size:255;"`
	Reamrk    string    `json:"reamrk" form:"reamrk" gorm:"column:reamrk;comment:;size:500;"`
	CreatedBy uint      `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint      `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint      `gorm:"column:deleted_by;comment:删除者"`
}

type ApiCommunityUser struct {
	ID        uint      `gorm:"primarykey"` // 主键ID
	UUID      uuid.UUID `json:"uuid" form:"uuid" gorm:"column:uuid;comment:;size:255;"`
	Nickname  string    `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:60;"`
	UserName  string    `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:60;"`
	Birthday  string    `json:"birthday" form:"birthday" gorm:"type:date;column:birthday;comment:;"`
	School    string    `json:"school" form:"school" gorm:"column:school;comment:;"`
	Education uint8     `json:"education" form:"education" gorm:"column:education;comment:;"`
	Status    uint8     `json:"status" form:"status" gorm:"column:status;comment:状态：1正常2禁止登录3禁言;"`
	Major     string    `json:"major" form:"major" gorm:"column:major;comment:;"`
	Avatar    string    `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;size:500;"`
	Phone     string    `json:"phone" form:"phone" gorm:"column:phone;comment:;size:20;"`
	Email     string    `json:"email" form:"email" gorm:"column:email;comment:;size:255;"`
	Reamrk    string    `json:"reamrk" form:"reamrk" gorm:"column:reamrk;comment:;size:500;"`
}

// TableName CommunityUser 表名
func (CommunityUser) TableName() string {
	return "community_user"
}
