package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/haithanh079/go-leaderboard/model"
	"github.com/haithanh079/go-leaderboard/model/response"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type LeaderboardController struct {
}


// HandleAddUser godoc
// swagger:operation POST /learderboard/add LearderBoard AddNewUser
//
// Add another user
//
// Add another user
//
// ---
// consumes:
// - application/x-www-form-urlencoded
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
		var Leaderboard []model.LeaderboardMember
		redisO, _ := c.Get("redis")
		redisClient := redisO.(*redis.Client)

		result, err := redisClient.Get("LEADERBOARD").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
		err = json.Unmarshal([]byte(result), &Leaderboard)
		var user = model.User{Name: c.PostForm("username"), Score: score}
		var currentUser = model.LeaderboardMember{user, calculateRank(user, redisClient)}
		Leaderboard = append(Leaderboard, currentUser)
		resetRank(Leaderboard, redisClient)

		serialized,err := json.Marshal(Leaderboard)
		redisClient.Set("LEADERBOARD", string(serialized), 1*time.Hour)

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
// consumes:
// - application/x-www-form-urlencoded
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
	var Leaderboard []model.LeaderboardMember
	redisO, _ := c.Get("redis")
	redisClient := redisO.(*redis.Client)

	result, err := redisClient.Get("LEADERBOARD").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	err = json.Unmarshal([]byte(result), &Leaderboard)
	responseDTO := response.LeaderboardRepsonse{response.BaseResponse{Success:true, Code:0, Msg:"Success"}, model.LearderBoard{Leaderboard}}
	c.JSON(http.StatusOK, responseDTO)
}

func calculateRank(user model.User, client *redis.Client) int {
	var ScoreList []int
	result, err := client.Get("SCORELIST").Result()
	if err != nil{
		panic(err)
	}
	err = json.Unmarshal([]byte(result), &ScoreList)
	var currentRank = getIndexIfExistScoreInArray(user.Score, ScoreList)
	if currentRank != -1{
		return currentRank
	}
	ScoreList = append(ScoreList, user.Score)
	sort.Ints(ScoreList)
	serialized,err := json.Marshal(ScoreList)
	client.Set("SCORELIST", string(serialized), 1*time.Hour)
	return getIndexIfExistScoreInArray(user.Score, ScoreList)
}

func resetRank(Leaderboard []model.LeaderboardMember, client *redis.Client)  {
	for index := 0; index < len(Leaderboard); index++ {
		currentUser := &Leaderboard[index]
		currentUser.Rank = calculateRank(currentUser.User, client)
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
