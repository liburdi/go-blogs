package controllers

import (
	"fmt"
	"golangschool/common/tools"
	"net/http"
	"time"
)

/**
 * err处理
 * @params error errMsg
 */
func checkErr(errMsg error) {
	if errMsg != nil {
		tools.WriteWithIoutil("./error.log", fmt.Sprintf("time:[%s] error:[%s] ", time.Now(), errMsg.Error()))
		return
	}
}


/**
 * 私有HTTP-ERROR处理
 * @params error errMsg
 * @params http.ResponseWriter w
 */
func checkHttpErr(errMsg error, w http.ResponseWriter) {
	if errMsg != nil {
		http.Error(w, errMsg.Error(), http.StatusInternalServerError)
		return
	}
}
