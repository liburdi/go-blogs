package controllers

import (
	"blog/config"
	"html/template"

		"io"
		"io/ioutil"
	//	"log"
	"net/http"
	//	"os"
	//	"syscall"
	"fmt"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/upload.html")
		checkErr(err)
		t.Execute(w, nil)
		return
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		checkErr(err)
		filename := h.Filename
		defer f.close()
		t, err := os.Create(config.UPLOAD_DIR + "/" + filename)
		checkErr(err)
		defer t.close()
		if _,err:=io.C
	}

}
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(1)
}
func ListHandler(w http.ResponseWriter, r *http.Request) {
	locals := make(map[string]interface{})
	locals["name"] = "asdasd"
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/list.html")
		checkErr(err)
		t.Execute(w, locals)
		return
	}
}
