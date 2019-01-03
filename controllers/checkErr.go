package controllers

import (
	"blog/common/tools"
	"net/http"
)


/**
 *私有ERROR处理
 */
func checkErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
		tools.WriteWithIoutil("./error.log",errMasg.Error())
		return
	}
}

/**
 *公有ERROR处理
 */
func CheckErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
		return
	}

}

/**
 *私有HTTP-ERROR处理
 */
func checkHttpErr(errMsg error,w http.ResponseWriter){
	if errMsg!=nil{
		http.Error(w,errMsg.Error(),http.StatusInternalServerError)
		return
	}
}


