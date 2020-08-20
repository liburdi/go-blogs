package controllers

import (
	initCommon "blog/common/init"
	. "blog/db"
	"blog/models"
	"net/http"
	"strconv"
);

/**
 * 测评列表
 * {
 *  type:1
 *
 *
 * }
 */
func ReviewList(response http.ResponseWriter, request *http.Request) {
	if request.Method == "GET" {
		println(request.Method)
	}
	reviewType := request.FormValue("type")
	page ,_:=strconv.Atoi(request.FormValue("page"))
	pageSize,_ :=strconv.Atoi(request.FormValue("pageSize"))
	review := make([]models.MentalityReview,0)
	err:= MasterDB.Where("type=?",reviewType).Limit(pageSize,(page-1)*pageSize).Find(&review)
	checkErr(err)
	//for _,v:=range review{
	//	switch v {
	//		case 1：
	//	}
	//}
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    review,
	}
	_ = api.ApiRestful(response)
}


func ModifyReview(response http.ResponseWriter, request *http.Request) {
	review := &models.MentalityReview{}

	//buf := make([]byte, 1024)
	//_ ,err :=request.Body.Read(buf)
	//checkErr(err)
	//body, _ := ioutil.ReadAll(request.Body)
	//	//err:=json.Unmarshal(body,review)

	review.Name=request.FormValue("name")
	review.Desc=request.FormValue("desc")

	Type, err:=strconv.ParseInt(request.FormValue("type"),10,8)
	review.Type=int8(Type)

	review.QuestionIds=request.FormValue("question_ids")

	isNice, err:=strconv.ParseInt(request.FormValue("is_nice"),10,8)
	review.IsNice=int8(isNice)

	_, err = MasterDB.Insert(review)
	checkErr(err)
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    review,
	}
	_ = api.ApiRestful(response)
}

/**
 * 问题列表
 */
func questionList(response http.ResponseWriter, request *http.Request) {

}
