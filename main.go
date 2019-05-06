package main

import (
	"github.com/haithanh079/go-leaderboard/routers"
)

func main() {
	r := routers.Router{}
	r.Init()
	r.Start()
}

