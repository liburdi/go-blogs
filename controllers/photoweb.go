package controllers

import (
	"blog/config"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

/**
 * 上传
 */
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles(config.TemplateDir + "/UploadTest/upload.html")
		checkErr(err)
		err = t.Execute(w, nil)
		checkHttpErr(err, w)
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
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}

}

/**
 * 浏览
 */
func ViewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := config.UPLOAD_DIR + "/" + imageId
	w.Header().Set("Content-Type", "image")
	http.ServeFile(w, r, imagePath)
}

/**
 * 列表
 */
func ListHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir("./uploads")
	checkHttpErr(err, w)
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
	checkHttpErr(err, w)
	err = t.Execute(w, locals)
	checkHttpErr(err, w)
}
