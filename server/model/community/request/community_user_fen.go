package request

// EditUserFenReq 关注
type EditUserFenReq struct {
	UserId    uint `json:"user_id"`     //
	FenUserId uint `json:"fen_user_id"` //
	IsFen     uint `json:"is_fen"`      // 关注？1是2取消
}
