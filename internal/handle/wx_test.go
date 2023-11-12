package handle

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"fmt"
	"sort"
	"testing"

	"WxGPT/internal/model"
)

func TestSha(t *testing.T) {
	list := []any{"token", "timestamp", "nonce"}
	sort.Slice(list, func(i, j int) bool {
		s1 := list[i].(string)
		s2 := list[j].(string)
		return s1 < s2
	})
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprint(list...)))
	fmt.Println(hex.EncodeToString(hash.Sum(nil)))
}

func TestMap(t *testing.T) {
	bytes := []byte(`<xml>
  <ToUserName>adda</ToUserName>
  <FromUserName>rewfe</FromUserName>
  <CreateTime>1348831860</CreateTime>
  <MsgType>jnvdhs</MsgType>
  <Content>adujwdn</Content>
  <MsgId>1234567890123456</MsgId>
  <MsgDataId>323</MsgDataId>
  <Idx>56654</Idx>
</xml>`)
	data := &model.ReceiveMsg{}
	_ = xml.Unmarshal(bytes, data)
	fmt.Println(data)
}
