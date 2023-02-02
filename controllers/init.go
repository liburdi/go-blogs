package controllers

import (
	"net/http"

	"github.com/liburdi/go-blogs/config"
	"github.com/liburdi/go-blogs/models"
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
}

func ParentController(w http.ResponseWriter, r *http.Request) {

}
