package router

import (
	"fmt"
	"golangschool/controllers"
	"golangschool/models"
	"net/http"
)

type KenRouter struct {
	router map[string]map[string]http.HandlerFunc
}

type Map struct {
	Method string
	Path   string
	Fn     func(http.ResponseWriter, *http.Request)
}

var AuthRouters = []string{
	"/pushProjectComment",
}
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
		Path: "/isOnline",
		Fn:   controllers.IsOnline,
	},
	{
		Path: "/authorDynamic",
		Fn:   controllers.AuthorDynamic,
	},
	{
		Method: "Get",
		Path:   "/niceComment",
		Fn:     controllers.NiceComment,
	},
	{
		Method: "Get",
		Path:   "/projectList",
		Fn:     controllers.ProjectList,
	},
	{
		Method: "Post",
		Path:   "/createProject",
		Fn:     controllers.CreateProject,
	},
	{
		Method: "Get",
		Path:   "/getProjectDetail",
		Fn:     controllers.ProjectDetail,
	},
	{
		Method: "Post",
		Path:   "/pushProjectComment",
		Fn:     controllers.PushProjectComment,
	},
	{
		Method: "Get",
		Path:   "/projectDetail",
		Fn:     controllers.ProjectDetailDisPlay,
	},
	{
		Method: "Get",
		Path:   "/projectIndex",
		Fn:     controllers.ProjectIndexDisPlay,
	},
	{
		Path: "/login/",
		Fn:   controllers.Login,
	},
	{
		Method: "Get",
		Path:   "/list",
		Fn:     controllers.ListHandler,
	},
	{
		Method: "Get",
		Path:   "/view",
		Fn:     controllers.ViewHandler,
	},
	{
		Method: "Post",
		Path:   "/upload",
		Fn:     controllers.UploadHandler,
	},
	{
		Method: "Get",
		Path:   "/reviewList",
		Fn:     controllers.ReviewList,
	},
	{
		Method: "Post",
		Path:   "/modifyReview",
		Fn:     controllers.ModifyReview,
	},
}

func (k *KenRouter) Routes() {
	for i := 0; i < len(Maps); i++ {
		cRoute := Maps[i]
		if cRoute.Method == "Get" || cRoute.Method == "get" {
			k.Get(cRoute.Path, cRoute.Fn)

		} else {
			k.Post(cRoute.Path, cRoute.Fn)
		}
	}
}

func NewRouter() *KenRouter {
	return new(KenRouter)
}

func (k *KenRouter) Get(pattern string, handler http.HandlerFunc) {
	k.HandleFunc("GET", pattern, handler)
}

func (k *KenRouter) Post(pattern string, handler http.HandlerFunc) {
	k.HandleFunc("POST", pattern, handler)
}

func (k *KenRouter) HandleFunc(method, pattern string, handle http.HandlerFunc) {
	if method == "" {
		panic("Method can not be null!")
	}

	if pattern == "" {
		panic("Pattern can not be null!")
	}

	if _, ok := k.router[method][pattern]; ok {
		panic("Pattern Exists!")
	}

	if k.router == nil {
		k.router = make(map[string]map[string]http.HandlerFunc)
	}

	if k.router[method] == nil {
		k.router[method] = make(map[string]http.HandlerFunc)
	}
	k.router[method][pattern] = handle
}

func (k *KenRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	status := models.Auth(w, r)
	fmt.Println(status)
	//权限检验
	if status == 1 {
		controllers.TemplateConfig["User"] = *models.BasicUserInfo
		fmt.Println("TempLateConfig:")
		fmt.Println(controllers.TemplateConfig)
	} else {
		for i := 0; i < len(AuthRouters); i++ {
			if AuthRouters[i] == r.URL.Path {
				fmt.Println("No Auth!!!")
				http.Error(w, "No Auth!!!", http.StatusUnauthorized)
				return
			}
		}
	}
	fmt.Println(status)
	fmt.Println(r.URL.String())
	fmt.Println(r.URL.Path)
	if f, ok := k.router[r.Method][r.URL.Path]; ok {
		f.ServeHTTP(w, r)
	} else {
	}
}
