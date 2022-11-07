package main

import (
	"example.com/model"
	"example.com/router"
	"github.com/gin-gonic/gin"
)

func handleRouter() *gin.Engine {
	r := router.RouterConfig()
	router.CreateRouter(r)
	return r
}

func main() {
	model.Connect()
	r := handleRouter()
	r.Run(":8080")
}
