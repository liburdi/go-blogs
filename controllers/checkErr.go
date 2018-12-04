package controllers

import "net/http"

func checkErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
		return
	}

}

func CheckErr(errMasg error) {
	if errMasg != nil {
		panic(errMasg)
		return
	}

}

func checkHttpErr(errMsg error,w http.ResponseWriter){
	if errMsg!=nil{
		http.Error(w,errMsg.Error(),http.StatusInternalServerError)
		return
	}
}
