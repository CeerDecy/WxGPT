package handle

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"testing"
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
	bytes := []byte(`{"msg":"success"}`)
	data := make(map[string]any)
	_ = json.Unmarshal(bytes, &data)
	fmt.Println(data)
}
