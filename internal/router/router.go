package router

import (
	"github.com/gin-gonic/gin"

	"WxGPT/internal/handle"
)

func Engine() *gin.Engine {
	router := gin.Default()
	router.GET("/wx", handle.Wx)
	router.POST("/wx", handle.Msg)
	return router
}
