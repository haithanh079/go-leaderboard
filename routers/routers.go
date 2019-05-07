package routers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/haithanh079/go-leaderboard/controller"
	"github.com/haithanh079/go-leaderboard/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type Router struct {
	Engine *gin.Engine
}

func (r *Router) ServeHTTP(http.ResponseWriter, *http.Request) {
}

func (r *Router) Init(testing bool) {
	if testing {
		if err := godotenv.Load("../.env"); err != nil {
			log.Println(err)
			panic(err)
		}
	}else {
		if err := godotenv.Load(); err != nil {
			log.Println(err)
			panic(err)
		}
	}

	e := gin.Default()

	isLocal := os.Getenv("LOCAL")
	hostName := "127.0.0.1"
	if strings.Compare(isLocal, "1") != 0 {
		hostName = os.Getenv("REDIS_HOSTNAME")
	}

	//Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", hostName),
		Password: os.Getenv("REDIS_PWD"),
		DB:       1,
	})

	log.Println("redis name: ", hostName)
	log.Println("redis pwd: ", os.Getenv("REDIS_PWD"))
	if _, err := redisClient.Ping().Result(); err != nil {
		panic(err)
	}
	var Leaderboard []model.LeaderboardMember
	serialized,err := json.Marshal(Leaderboard)
	if err != nil {
		panic(err)
	}
	redisClient.Set("LEADERBOARD", string(serialized), 1*time.Hour)
	var ScoreList []int
	serialized, err = json.Marshal(ScoreList)
	if err != nil {
		panic(err)
	}
	redisClient.Set("SCORELIST", string(serialized), 1*time.Hour)

	e.Use(func(c *gin.Context) {
		c.Set("redis", redisClient)
	})

	user := e.Group("/leaderboard")
	{
		leaderboard := controller.LeaderboardController{}
		user.POST("/add", leaderboard.AddUserToLeaderboard)
		user.GET("/get", leaderboard.GetLeaderBoard)
	}
	r.Engine = e
}

//Start router
func (r *Router) Start() {
	log.Fatalln(r.Engine.Run(":8000"))
}
