package controllers

import (
	"blog/common/tools"
	"fmt"
	"net/http"
	"time"
)

/**
 *私有ERROR处理
 */
func checkErr(errMsg error) {
	if errMsg != nil {
		tools.WriteWithIoutil("./error.log", fmt.Sprintf("time:[%s] error:[%s] ",time.Now(),errMsg.Error()))
		return
	}
}

/**
 *公有ERROR处理
 */
func CheckErr(errMsg error) {
	if errMsg != nil {
		panic(errMsg)
		return
	}
}

/**
 *私有HTTP-ERROR处理
 */
func checkHttpErr(errMsg error, w http.ResponseWriter) {
	if errMsg != nil {
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
		return
	}
}
