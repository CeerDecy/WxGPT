package router

import (
	"github.com/gin-gonic/gin"

	"WxGPT/internal/gpt/gptclient"
	"WxGPT/internal/handle"
	"WxGPT/internal/session"
)

func Engine() *gin.Engine {
	router := gin.Default()
	router.GET("/wx", handle.Auth)
	router.POST("/wx", handle.ReceiveAndReturn)
	router.GET("/stream", handle.StreamWeb)
	router.GET("/ask", func(ctx *gin.Context) {
		query, _ := ctx.GetQuery("q")
		client := gptclient.DefaultClient()
		stream, err := client.GetStreamResponse(query)
		if err != nil {
			ctx.String(200, err.Error())
			return
		}
		sess := session.NewSession(stream)
		session.ChatSession.Set("1", sess)
		sess.ReadStream()
		ctx.String(200, "http://localhost:80/stream?sid=1")
	})
	return router
}
