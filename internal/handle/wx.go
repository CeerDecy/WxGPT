package handle

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"

	"github.com/gin-gonic/gin"

	"WxGPT/internal/model"
	"WxGPT/internal/tools"
)

func Wx(ctx *gin.Context) {
	signature, _ := ctx.GetQuery("signature")
	timestamp, _ := ctx.GetQuery("timestamp")
	//echostr, _ := ctx.GetQuery("echostr")
	nonce, _ := ctx.GetQuery("nonce")
	openid, _ := ctx.GetQuery("openid")
	token := "WxGPT"
	if tools.Auth(signature, timestamp, nonce, token) {

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
		ctx.XML(200, model.DefaultTextResp(openid, token, "msg received : "+data.Content))
	} else {
		ctx.XML(200, model.DefaultTextResp(openid, token, "认证失败"))
	}

}
