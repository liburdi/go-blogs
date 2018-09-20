package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	h := md5.New()
	h.Write([]byte("123")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	pas := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	fmt.Println(pas)

}
