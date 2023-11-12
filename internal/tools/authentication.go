package tools

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"sort"
)

func Auth(signature string, timestamp string, nonce string, token string) bool {
	list := []any{token, timestamp, nonce}
	sort.Slice(list, func(i, j int) bool {
		s1 := list[i].(string)
		s2 := list[j].(string)
		return s1 < s2
	})
	hash := sha1.New()
	hash.Write([]byte(fmt.Sprint(list...)))
	encodeToString := hex.EncodeToString(hash.Sum(nil))
	return encodeToString == signature
}
