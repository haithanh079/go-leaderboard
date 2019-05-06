package model

type LeaderboardMember struct {
	User User `json:"user"`
	Rank int `json:"rank"`
}
