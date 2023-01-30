package main

import (
	"github.com/gin-gonic/gin"
	"main.go/src/service"
)

func main() {
	r := gin.Default()
	r.GET("/", func(context *gin.Context) {
		context.JSON(200, "hello")
	})
	r.GET("/namespace", service.ListNamespace)
	r.GET("/deployments", service.ListDeployment)
	r.GET("/service", service.ListService)
	r.GET("/pods", service.ListAllPod)
	r.Run()
}
