package route

import (
	"https/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.CheckSig)
	router.POST("/", controller.HandleTextMsg)
	return router
}
