package request

// EditUserCollectReq 收藏
type EditUserCollectReq struct {
	UserId    uint `json:"user_id"`    //
	PostId    uint `json:"post_id"`    //
	IsCollect uint `json:"is_collect"` // 收藏？1是2取消
}
