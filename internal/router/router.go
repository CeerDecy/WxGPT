package router

import (
	"github.com/gin-gonic/gin"

	"WxGPT/internal/handle"
)

func Engine() *gin.Engine {
	router := gin.Default()
	router.GET("/wx", handle.Auth)
	router.POST("/wx", handle.ReceiveAndReturn)
	return router
}
