package tools

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(v []byte) string {
	m := md5.New()
	m.Write(v)
	return hex.EncodeToString(m.Sum(nil))
}
