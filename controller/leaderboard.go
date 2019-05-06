package controller

import (
	"sort"
	"go-leaderboard/model"
)

type LeaderboardMember model.LeaderboardMember
type User model.User

var Leaderboard []LeaderboardMember
var ScoreList []int;

func addUserToLeaderboard(user User)  {
	Leaderboard = append(Leaderboard, LeaderboardMember{user, calculateRank(user)})
}

func calculateRank(user User) int {
	var currentRank = getIndexIfExistInArray(user.Score, ScoreList)
	if currentRank != -1{
		return currentRank
	}
	ScoreList = append(ScoreList, user.Score)
	sort.Ints(ScoreList)
	return getIndexIfExistInArray(user.Score, ScoreList)
}

func getIndexIfExistInArray(num int, list []int) int {
	for currentIndex, currentNum := range list {
		if currentNum == num {
			return currentIndex
		}
	}
	return -1
}