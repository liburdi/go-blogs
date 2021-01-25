package controllers

import (
	"golangschool/models"
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
/**
 * @params http.ResponseWrite w
 * @params *http.Request r
 */
func ParentController(w http.ResponseWriter, r *http.Request){
	status:=models.Auth(w,r)
	fmt.Println(status)
}
