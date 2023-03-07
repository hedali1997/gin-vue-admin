package request

// EditUserBlockReq 收藏请求
type EditUserBlockReq struct {
	UserId      uint `json:"user_id"`       //
	BlockUserId uint `json:"block_user_id"` //
	IsBlock     uint `json:"is_block"`      // 拉黑？1是2取消
}
