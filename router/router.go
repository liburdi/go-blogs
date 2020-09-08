package router

import (
	"../controllers"
	"net/http"
)

// Map 路由
type Map struct {
	Path string
	Fn   func(http.ResponseWriter, *http.Request)
}

// Maps 路由列表
var Maps = []*Map{

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
		Path: "/note",
		Fn:   controllers.Note,
	},
	{
		Path: "/is_online",
		Fn:   controllers.IsOnline,
	},
	{
		Path: "/author_dynamic",
		Fn:   controllers.AuthorDynamic,
	},
	{
		Path: "/nice_comment",
		Fn:   controllers.NiceComment,
	},
	{
		Path: "/login/",
		Fn:   controllers.Login,
	},
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
	{
		Path: "/reviewList",
		Fn:   controllers.ReviewList,
	},
	{
		Path: "/modifyReview",
		Fn:   controllers.ModifyReview,
	},

}

// Routes 操作
func Routes() {
	for i := 0; i < len(Maps); i++ {
		cRoute := Maps[i]
		http.HandleFunc(cRoute.Path, cRoute.Fn)
	}
}
