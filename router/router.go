package router

import (
	"github.com/gin-gonic/gin"

	"WxGPT/handle"
)

func Engine() *gin.Engine {
	router := gin.Default()
	router.Any("/wx", handle.Wx)
	return router
}
