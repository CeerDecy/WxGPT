package handle

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	"WxGPT/internal/model"
	"WxGPT/internal/tools"
)

func Msg(ctx *gin.Context) {
	bytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println(err)
	}
	var data model.TextReceive
	err = xml.Unmarshal(bytes, &data)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(data)
	//response := model.TextResponse{
	//	ToUserName:   data.FromUserName,
	//	FromUserName: data.ToUserName,
	//	CreateTime:   uint64(time.Now().Unix()),
	//	MsgType:      data.MsgType,
	//	Content:      "receive msg :" + data.Content,
	//}
	ctx.Data(
		200,
		"application/xml",
		[]byte(fmt.Sprintf(`<xml><ToUserName><![CDATA[%s]]></ToUserName><FromUserName><![CDATA[%s]]></FromUserName><CreateTime>%d</CreateTime><MsgType><![CDATA[text]]></MsgType><Content><![CDATA[%s]]></Content></xml>`,
			data.FromUserName,
			data.ToUserName,
			uint64(time.Now().Unix()),
			"receive msg :"+data.Content)))
	//ctx.XML(200, response)
}

func Wx(ctx *gin.Context) {
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
