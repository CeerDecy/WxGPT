package router

import (
	"github.com/gin-gonic/gin"
)

func Engine() *gin.Engine {
	router := gin.Default()
	return router
}
