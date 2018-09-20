package router

import (
	"blog/controllers"
	"net/http"
)

// RouterMap 路由
type RouterMap struct {
	Path string
	Fn   func(http.ResponseWriter, *http.Request)
}

// RouterMaps 路由列表
var RouterMaps = []*RouterMap{
	{
		Path: "/item/",
		Fn:   controllers.ArticleItem,
	},
	{
		Path: "/login/",
		Fn:   controllers.Login,
	},
	{
		Path: "/test/",
		Fn:   controllers.Test,
	},
	{
		Path: "/login1/",
		Fn:   controllers.Login1,
	},
}

// Routes 操作
func Routes() {
	for i := 0; i < len(RouterMaps); i++ {
		cRoute := RouterMaps[i]
		http.HandleFunc(cRoute.Path, cRoute.Fn)
	}
}
