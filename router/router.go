package router

import (
	"github.com/liburdi/go-blogs/controllers"
	"github.com/liburdi/go-blogs/models"
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
		Method: "Get",
		Path:   "/",
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
		Path:   "/upload",
		Fn:     controllers.UploadHandler,
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
	//权限检验
	if status == nil {
		controllers.TemplateConfig["User"] = *models.BasicUserInfo
	} else {
		for i := 0; i < len(AuthRouters); i++ {
			if AuthRouters[i] == r.URL.Path {
				//TODO
			}
		}
	}
	if f, ok := k.router[r.Method][r.URL.Path]; ok {
		f.ServeHTTP(w, r)
	}
	return
}
