package community

type UserBlock struct {
	ID          uint `gorm:"primarykey"` // 主键ID
	UserId      uint `json:"user_id" gorm:"column:user_id;"`
	BlockUserId uint `json:"block_user_id" gorm:"column:block_user_id;"`
}
