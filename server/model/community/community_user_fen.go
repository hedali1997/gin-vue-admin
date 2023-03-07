package community

type UserFen struct {
	ID        uint `gorm:"primarykey"` // 主键ID
	UserId    uint `json:"user_id" gorm:"column:user_id;"`
	FenUserId uint `json:"fen_user_id" gorm:"column:fen_user_id;"`
}
