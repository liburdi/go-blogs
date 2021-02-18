package controllers

import (
	"fmt"
	"golangschool/config"
	"golangschool/models"
	"net/http"
)

type UserInfo struct {
	userId   int
	userName string
}
var TemplateConfig = map[string]interface{}{
	"JsPath":        config.JsPath,
	"BootStrapPath": config.BootStrapPath,
	"ImgPath":       config.ImgPath,
	"CssPath":       config.CssPath,
	"StaticDir":     config.StaticDir,
	"User":          models.BasicUserInfo,
}
func init() {
	fmt.Println("init controller")
}

/**
 * @params http.ResponseWrite w
 * @params *http.Request r
 */
func ParentController(w http.ResponseWriter, r *http.Request) {

}
