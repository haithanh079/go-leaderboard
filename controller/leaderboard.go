package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	//"encoding/json"
	"github.com/haithanh079/go-leaderboard/model"
	"github.com/haithanh079/go-leaderboard/model/response"
	"sort"
	//"github.com/gomodule/redigo/redis"
)

type LeaderboardController struct {
}


var Leaderboard []model.LeaderboardMember
var ScoreList []int


/**
Add user & get learderboard
 */

// HandleAddUser godoc
// swagger:operation POST /learderboard/add LearderBoard AddNewUser
//
// Add another user
//
// Add another user
//
// ---
// produces:
// - application/json
// schemes:
// - https
// security:
// - api_key:
// parameters:
// - name: username
//   in: body
//   description: username
//   required: true
//   type: string
//   example: user
// - name: score
//   in: body
//   description: score
//   required: true
//   type: int
//   example: 1
// responses:
//   '200':
//     description: OK
//     schema:
//       type: object
//       "$ref": "#/definitions/MemberResponse"
//   '500':
//     description: Internal Server
//     schema:
//       type: object
//       "$ref": "#/definitions/MemberResponse"
func (LeaderboardController *LeaderboardController) AddUserToLeaderboard(c *gin.Context) {
	responseDTO := response.MemberResponse{}
	score, err := strconv.Atoi(c.PostForm("score"))
	if err != nil {
		responseDTO = response.MemberResponse{response.BaseResponse{Success:false, Code: 1, Msg: "Score must be Integer!"}, model.LeaderboardMember{}}
	}else {
		var user = model.User{Name: c.PostForm("username"), Score: score}
		var currentUser = model.LeaderboardMember{user, calculateRank(user)}
		Leaderboard = append(Leaderboard, currentUser)
		resetRank()
		responseDTO.Success = true
		responseDTO.Code = 0
		responseDTO.Msg = "Success"
		responseDTO.Data = currentUser
	}
	c.JSON(http.StatusOK, responseDTO)
}

// HandleAddUser godoc
// swagger:operation GET /learderboard/get LearderBoard GetLeaderboard
//
// Get leaderboard
//
// Get leaderboard
//
// ---
// produces:
// - application/json
// schemes:
// - https
// security:
// - api_key:
// parameters:
// responses:
//   '200':
//     description: OK
//     schema:
//       type: object
//       "$ref": "#/definitions/LeaderboardResponse"
//   '500':
//     description: Internal Server
//     schema:
//       type: object
//       "$ref": "#/definitions/LeaderboardResponse"
func (LeaderboardController *LeaderboardController) GetLeaderBoard(c *gin.Context)  {
	responseDTO := response.LeaderboardRepsonse{response.BaseResponse{Success:true, Code:0, Msg:"Success"}, model.LearderBoard{Leaderboard}}
	c.JSON(http.StatusOK, responseDTO)
}

func calculateRank(user model.User) int {
	var currentRank = getIndexIfExistScoreInArray(user.Score, ScoreList)
	if currentRank != -1{
		return currentRank
	}
	ScoreList = append(ScoreList, user.Score)
	sort.Ints(ScoreList)
	return getIndexIfExistScoreInArray(user.Score, ScoreList)
}

func resetRank()  {
	for index := 0; index < len(Leaderboard); index++ {
		currentUser := &Leaderboard[index]
		currentUser.Rank = calculateRank(currentUser.User)
	}
}

func getIndexIfExistScoreInArray(num int, list []int) int {
	for currentIndex, currentNum := range list {
		if currentNum == num {
			return currentIndex
		}
	}
	return -1
}