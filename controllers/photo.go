package controllers

import (
	"fmt"
	"github.com/liburdi/go-blogs/config"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/upload.html")
		if err != nil {
			return
		}
		err = t.Execute(w, nil)
		if err != nil {
			return
		}
	}

	if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		checkErr(err)
		filename := h.Filename
		defer checkErr(f.Close())
		t, err := os.Create(config.UploadDir + "/" + filename)
		checkErr(err)

		if _, err := io.Copy(t, f); err != nil {
			checkErr(err)
		}
		defer checkErr(t.Close())
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := config.UploadDir + "/" + imageId
	fmt.Printf(imagePath)
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	checkHttpErr(err, w)
	var locals = make(map[string]interface{})
	var images = []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/list.html")
	checkHttpErr(err, w)
	err = t.Execute(w, locals)
	checkHttpErr(err, w)
}
