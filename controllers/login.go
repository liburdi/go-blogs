package controllers

import (
	"golangschool/config"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		locals := make(map[string]interface{})
		locals["name"] = "asdasd"
		t, _ := template.ParseFiles(config.TemplateDir+"/Index/login.html", config.HeaderPath, config.FooterPath)
		err := t.Execute(w, locals)
		if err != nil {
			fmt.Println("err")
		}
	}
	if r.Method == "POST" {
		v := r.FormValue("username")
		fmt.Println(v)
	}

}
