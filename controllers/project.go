package controllers

import (
	"fmt"
	initCommon "golangschool/common/init"
	. "golangschool/db"
	"golangschool/models"
	"html/template"
	"net/http"
	"strconv"
)

/**
 * api result
 */
type projectListResult struct {
	*models.Php41Project
	Author string `json:"author"`
}

/**
 * 开源项目列表
 */
func ProjectList(w http.ResponseWriter, r *http.Request) {
	projectList := make([]*models.Php41Project, 0)
	err := MasterDB.Where("is_del=0").Find(&projectList)
	checkHttpErr(err, w)
	projectListResult := make([]projectListResult, len(projectList))
	user := make([]*models.Php41Users, 0)
	var temp *models.Php41Project
	for k, v := range projectList {
		err = MasterDB.Where("user_id = ?", v.CreatorId).Find(&user)
		fmt.Println(user)
		if len(user) < len(projectList) {
			user = append(user, &models.Php41Users{Username: "xx"})
		}
		temp = v
		projectListResult[k].Php41Project = temp
		projectListResult[k].Author = user[k].Username
		fmt.Println(projectListResult)

	}
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    projectListResult,
	}

	api.ApiRestful(w)

}

/**
 * 发布项目
 */
func CreateProject(w http.ResponseWriter, r *http.Request) {
	Project := new(models.Php41Project)
	Project.Title = r.FormValue("title")
	Project.Desc = r.FormValue("desc")
	Project.OriginUrl = r.FormValue("origin_url")
	if Project.Title == "" || Project.Desc == "" || Project.OriginUrl == "" {
		initCommon.Error(3000, w)
		return
	}
	affected, err := MasterDB.Insert(Project)
	if affected > 0 && err == nil {
		initCommon.Success(w)
	} else {
		initCommon.Error(3000, w)
	}
	fmt.Println(err)

}

func ProjectDetail(w http.ResponseWriter, r *http.Request) {
	project := new(models.Php41Project)
	project.Id, _ = strconv.Atoi(r.FormValue("id"))
	has, err := MasterDB.Get(project)
	checkHttpErr(err, w)
	if has {
		api := initCommon.ApiRestful{
			Code:    200,
			Message: "Success",
			Data:    project,
		}

		api.ApiRestful(w)
	} else {
		initCommon.Error(2001, w)
	}

}

func PushProjectComment(w http.ResponseWriter, r *http.Request) {
	comment := new(models.Php41Ooxx)
	//comment.TargetId, _ = strconv.Atoi(r.FormValue("project_id"))
	comment.TargetId=26
	comment.Msg = r.FormValue("msg")
	comment.ToUser, _ = strconv.Atoi(r.FormValue("to_user"))
	comment.TargetType = 2
	comment.FromUser = models.BasicUserInfo.UserId
	if comment.TargetId == 0 || comment.Msg == ""  {
		initCommon.Error(3000, w)
		return
	}
	affected, err := MasterDB.Insert(comment)
	if affected > 0 && err == nil {
		initCommon.Success(w)
	} else {
		initCommon.Error(3000, w)
	}
	fmt.Println(err)

}

func ProjectDetailDisPlay(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/footer.html", "templates/project/detail.html"))
	daysOfWeek := map[string]interface{}{"pageTitle": "a", "config": TemplateConfig}
	checkHttpErr(t.ExecuteTemplate(w, "projectDetail", daysOfWeek), w)
}

func ProjectIndexDisPlay(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/layout.html", "templates/footer.html", "templates/project/index.html"))
	daysOfWeek := map[string]interface{}{"pageTitle": "a", "config": TemplateConfig}
	checkHttpErr(t.ExecuteTemplate(w, "projectIndex", daysOfWeek), w)
}
