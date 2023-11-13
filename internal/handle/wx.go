package handle

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"

	"WxGPT/internal/gpt/gptclient"
	"WxGPT/internal/model"
	"WxGPT/internal/session"
	"WxGPT/internal/tools"
)

func ReceiveAndReturn(ctx *gin.Context) {
	signature, _ := ctx.GetQuery("signature")
	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
	}
	var data model.TextReceive
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		log.Println(err)
	}
	log.Printf("[Unmarshal data ] : %+v", data)
	client := gptclient.DefaultClient()
	//response, err := client.GetResponse(data.Content)
	stream, err := client.GetStreamResponse(data.Content)
	if err != nil {
		if err.Error() == `Post "https://proxy.geekai.co/v1/chat/completions": context deadline exceeded` {
			ctx.Data(
				200,
				"application/xml",
				[]byte(model.DefaultTextResp(data.FromUserName, data.ToUserName, "AI响应超时")))
			return
		}
		ctx.Data(
			200,
			"application/xml",
			[]byte(model.DefaultTextResp(data.FromUserName, data.ToUserName, err.Error())))
		return
	}
	sid := tools.Md5([]byte(signature))
	sess := session.NewSession(stream)
	session.ChatSession.Set(sid, sess)
	sess.ReadStream()
	ctx.Data(
		200,
		"application/xml",
		[]byte(model.DefaultTextResp(
			data.FromUserName,
			data.ToUserName,
			fmt.Sprintf("由于GPT响应时间可能会比较长，获取结果请访问此链接 http://101.43.101.59/stream?sid=%s", sid))))

}

func Auth(ctx *gin.Context) {
	signature, _ := ctx.GetQuery("signature")
	timestamp, _ := ctx.GetQuery("timestamp")
	nonce, _ := ctx.GetQuery("nonce")
	token := "WxGPT"
	if tools.Auth(signature, timestamp, nonce, token) {
		ctx.String(200, signature)
	} else {
		ctx.String(200, "")
	}
}
