package response

import "github.com/haithanh079/go-leaderboard/model"

// swagger:model
type MemberResponse struct {
	BaseResponse
	Data model.LeaderboardMember `json:"data"`
} 
