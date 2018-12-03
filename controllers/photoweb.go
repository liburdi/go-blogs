package controllers

import (
	"blog/config"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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
		defer f.Close()
		t, err := os.Create(config.UPLOAD_DIR + "/" + filename)
		checkErr(err)
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			checkErr(err)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := config.UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return

	}
	//方式一
	//	var listHtml string
	//	for _, fileInfo := range fileInfoArr {
	//		imgid := fileInfo.Name()
	//		listHtml += "<li><a href=\"/view?id=" + imgid + "\">imgid</a></li>"

	//	}
	//	io.WriteString(w, "<div>"+listHtml+"</div>")
	//方式二
	var locals = make(map[string]interface{})
	var images = []string{}
	for _, fileInfo := range fileInfoArr {
		images = append(images, fileInfo.Name())
	}
	locals["images"] = images
	t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/list.html")
	if err != nil {
		http.Error(w, err.Error(),
			http.StatusInternalServerError)
		return
	}
	t.Execute(w, locals)
}
