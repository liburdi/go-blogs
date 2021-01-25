package main

import (
	"golangschool/router"
	"log"
	"net/http"
)

func main() {
	//路由注册
	router.Routes()
	//静态文件服务挂载
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//监听服务端口
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println("Server Start Failed")
	} else {
		log.Println("Listening on 0.0.0.0:8080")
	}
}
