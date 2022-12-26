package main

import (
	"part2/internal/consts"
	"part2/internal/model"
	"part2/internal/service"
	"time"
)

func main() {
	date, _ := time.Parse(consts.LAYOUT, "2000-04-28")
	date2 := time.Now()
	date2.Date()
	u := model.NewUser("alex", "alex228", "qwe", "89658481895", date)
	//fmt.Print(u)

	router := service.SetupRouter()

	router.Run("localhost:8080")
}
