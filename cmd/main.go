package main

import (
	"fmt"
	"part2/internal/consts"
	"part2/internal/model"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	date, _ := time.Parse(consts.LAYOUT, "2000-04-28")
	date2 := time.Now()
	date2.Date()
	u := model.NewUser("alex", "alex228", "qwe", "89658481895", date)
	fmt.Print(u)

	router := gin.Default()
	router.GET("/users", model.User)

	router.Run("localhost:8080")
}
