// 自动生成模板CommunityUser
package community

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// CommunityUser 结构体
type CommunityUser struct {
	global.GVA_MODEL
	Nickname  string     `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;size:60;"`
	UserName  string     `json:"userName" form:"userName" gorm:"column:user_name;comment:用户名;size:60;"`
	Password  string     `json:"password" form:"password" gorm:"column:password;comment:密码;size:255;"`
	Birthday  *time.Time `json:"birthday" form:"birthday" gorm:"column:birthday;comment:;"`
	School    string     `json:"school" form:"school" gorm:"column:school;comment:;"`
	Education *int       `json:"education" form:"education" gorm:"column:education;comment:;"`
	Major     string     `json:"major" form:"major" gorm:"column:major;comment:;"`
	Avatar    string     `json:"avatar" form:"avatar" gorm:"column:avatar;comment:;size:500;"`
	Phone     string     `json:"phone" form:"phone" gorm:"column:phone;comment:;size:20;"`
	Email     string     `json:"email" form:"email" gorm:"column:email;comment:;size:255;"`
	Reamrk    string     `json:"reamrk" form:"reamrk" gorm:"column:reamrk;comment:;size:500;"`
	CreatedBy uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CommunityUser 表名
func (CommunityUser) TableName() string {
	return "community_user"
}
