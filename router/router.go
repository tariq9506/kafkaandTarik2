package router

import (
	"tarik_consumer/kafkaandTarik2/controller"

	"github.com/gin-gonic/gin"
)

func AddRouter(router *gin.RouterGroup) {
	router.GET("/listen", controller.ListenForLocationUpdates)
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	AddRouter(&router.RouterGroup)
	return router
}
