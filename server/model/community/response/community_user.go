package response

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/community"
)

type UserResponse struct {
	User community.CommunityUser `json:"user"`
}

type LoginResponse struct {
	User      community.ApiCommunityUser `json:"user"`
	Token     string                     `json:"token"`
	ExpiresAt int64                      `json:"expiresAt"`
}

type RegisterResponse struct {
	User community.ApiCommunityUser `json:"user"`
}
