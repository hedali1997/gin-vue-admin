package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
	"time"
)

type CommunityUserSearch struct {
	community.CommunityUser
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

// Modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user id，避免越权
	Password    string `json:"password"`    // 密码
	NewPassword string `json:"newPassword"` // 新密码
}

// Register User register structure
type Register struct {
	UserName      string `json:"user_name" example:"用户名"`
	Phone         string `json:"phone" example:"手机号"`
	Code          string `json:"code" example:"短信验证码"`
	Password      string `json:"password" example:"密码"`
	CheckPassword string `json:"check_passWord" example:"确认密码"`
}

// User login structure
type PhoneLogin struct {
	Phone     string `json:"phone"`     // 用户名
	Password  string `json:"password"`  // 密码
	Captcha   string `json:"captcha"`   // 验证码
	CaptchaId string `json:"captchaId"` // 验证码ID
}

type PhoneCodeLogin struct {
	Phone string `json:"phone"` // 手机号
	Code  string `json:"code"`  // 短信验证码
}
