package service

import "github.com/gin-gonic/gin"

type Router struct {
	router gin.Engine
	//handler Handler
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/user/:login", func(ctx *gin.Context) {

	})
	r.POST("/user/:login", func(ctx *gin.Context) {

	})
	r.POST("/user/login", func(ctx *gin.Context) {

	})
	return r
}

func AddHandlerGet(endpoint string, handlerFunc gin.HandlerFunc) {

}

func (router *Router) Run(port string) {
	router.router.Run(port)
}
