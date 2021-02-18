package main

import (
	"golangschool/router"
	"log"
	"net/http"
)

func main() {
	//路由注册
	kenRouter := router.NewRouter()
	kenRouter.Routes()

	//静态文件服务挂载
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	//监听服务端口
	err := http.ListenAndServe(":8083", kenRouter)
	if err != nil {
		log.Println("Server Start Failed")
	} else {
		log.Println("Listening on 0.0.0.0:8083")
	}
}
