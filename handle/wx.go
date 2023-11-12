package handle

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Wx(ctx *gin.Context) {
	signature, _ := ctx.GetQuery("signature")
	log.Println("signature ", signature)
	timestamp, _ := ctx.GetQuery("timestamp")
	log.Println("timestamp ", timestamp)
	nonce, _ := ctx.GetQuery("nonce")
	log.Println("nonce ", nonce)
	ctx.JSON(200, gin.H{
		"msg":  "success",
		"code": 200,
	})
}
