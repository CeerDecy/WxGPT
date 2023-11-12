package handle

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"sort"

	"github.com/gin-gonic/gin"

	"WxGPT/internal/model"
)

func Wx(ctx *gin.Context) {
	signature, _ := ctx.GetQuery("signature")
	timestamp, _ := ctx.GetQuery("timestamp")
	echostr, _ := ctx.GetQuery("echostr")
	nonce, _ := ctx.GetQuery("nonce")
	token := "WxGPT"
	list := []any{token, timestamp, nonce}
	sort.Slice(list, func(i, j int) bool {
		s1 := list[i].(string)
		s2 := list[j].(string)
		return s1 < s2
	})
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprint(list...)))
	encodeToString := hex.EncodeToString(hash.Sum(nil))
	if encodeToString == signature {
		bytes, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			log.Println(err)
		}
		var data model.ReceiveMsg
		err = xml.Unmarshal(bytes, &data)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(data)
		ctx.String(200, echostr)
	} else {
		ctx.String(200, "")
	}
}
