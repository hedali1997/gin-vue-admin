package community

type UserCollectPost struct {
	ID     uint `gorm:"primarykey"` // 主键ID
	UserId uint `json:"user_id" gorm:"column:user_id;"`
	PostId uint `json:"post_id" gorm:"column:post_id;"`
}
