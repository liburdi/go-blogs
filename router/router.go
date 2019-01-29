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
		Path: "/test/",
		Fn:   controllers.Test,
	},
	{
		Path: "/testApi/",
		Fn:   controllers.TestApi,
	},
	{
		Path: "/testAuth/",
		Fn:   controllers.TestAuth,
	},
	{
		Path: "/home/index/note",
		Fn:   controllers.Note,
	},
	{
		Path: "/testconn/",
		Fn:   controllers.TestConn,
	},
	{
		Path: "/login/",
		Fn:   controllers.Login,
	},
	//upload image gdemo
	{
		Path: "/",
		Fn:   controllers.ListHandler,
	},
	{
		Path: "/view",
		Fn:   controllers.ViewHandler,
	},
	{
		Path: "/upload",
		Fn:   controllers.UploadHandler,
	},
}

// Routes 操作
func Routes() {
	for i := 0; i < len(RouterMaps); i++ {
		cRoute := RouterMaps[i]
		http.HandleFunc(cRoute.Path, cRoute.Fn)
	}
}
