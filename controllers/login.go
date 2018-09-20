package controllers

import (
	"blog/config"
	"blog/models"
	"blog/redis"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var usersModel models.UsersModel
	usersModel.SetTableName()
	nam := r.FormValue("username")
	pas := r.FormValue("passwd")
	var userId int64
	var userName string
	var passWd string
	selectField := models.SelectValues{
		"user_id":  &userId,
		"username": &userName,
		"password": &passWd,
	}
	h := md5.New()
	h.Write([]byte(pas)) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Printf("%s\n", hex.EncodeToString(cipherStr))      // 输出加密结果
	pas = fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果

	where := models.WhereValues{}
	where["user_id"] = models.WhereCondition{">", "1"}
	where["username"] = models.WhereCondition{"=", nam}
	where["password"] = models.WhereCondition{"=", pas}
	qr, err := usersModel.Query(selectField, where, 0, 1)
	checkErr(err)
	//	v := redis.Set("username", "chenzhaofei")
	//	if v != "ok" {
	//		fmt.Println(v)
	//	}
	//	v, err = redis.Get("username")
	//	checkErr(err)
	SessionUser := v
	t, _ := template.ParseFiles(config.TemplateDir+"/Index/login.html", config.HeaderPath, config.FooterPath)
	err = t.Execute(w, {
		
	})
	checkErr(err)
}
func Login1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(config.TemplateDir+"/Index/login.html", config.HeaderPath, config.FooterPath)
	Result := []models.SelectResult{map[string]interface{}{"1234": "asdasd"}, map[string]interface{}{"12344": "asdasd1"}}
	err := t.Execute(w, struct {
		SessionUser []models.SelectResult
	}{
		SessionUser: Result,
	})
	checkErr(err)
}
