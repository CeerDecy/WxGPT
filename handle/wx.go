package handle

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Wx(ctx *gin.Context) {
	header := ctx.GetHeader("signature")
	log.Println("header ", header)
	value, _ := ctx.Get("signature")
	log.Println("body ", value)
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
	})
}
