package controller

import (
	"crypto/md5"
	"encoding/hex"
)

//生成md5加密字符串
func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
