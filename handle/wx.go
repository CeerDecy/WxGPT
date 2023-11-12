package handle

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"
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
		ctx.String(200, echostr)
	} else {
		ctx.String(200, "")
	}
}
