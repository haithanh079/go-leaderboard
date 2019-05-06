package response

import "github.com/haithanh079/go-leaderboard/model"

// swagger:model
type LeaderboardRepsonse struct {
	BaseResponse
	Data model.LearderBoard `json:"data"`
}
