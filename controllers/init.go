package controllers

import (
	"blog/models"
	"fmt"
	"net/http"
)

type UserInfo struct {
	userId int
	userName string
}
func init(){
	fmt.Println("init controller")
}
func ParentChontroller(w http.ResponseWriter, r *http.Request){
	status:=models.Auth(w,r)
	fmt.Println(status)
}
