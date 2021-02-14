package main

import (
	"github.com/ecto0310/online_judge_backend/src/router"
)

func main() {
	r := router.InitRouter()
	r.Logger.Fatal(r.Start(":1323"))
}
