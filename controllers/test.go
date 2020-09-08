package controllers

import (
	initCommon "../common/init"
	"../common/tools"
	. "../db"
	"../models"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	h := md5.New()
	h.Write([]byte("123")) // 需要加密的字符串为 123456
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	pas := fmt.Sprintf("%s", hex.EncodeToString(cipherStr)) // 输出加密结果
	fmt.Println(pas)

}
func TestApi(w http.ResponseWriter, r *http.Request) {
	users := make([]*models.Php41Users, 0)
	err := MasterDB.Where("(user_id>?)", 0).Find(&users)
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    users,
	}
	err = api.ApiRestful(w)
	checkErr(err)
}
func TestAuth(w http.ResponseWriter, r *http.Request) {
	ParentChontroller(w, r)
}

func Note(w http.ResponseWriter, r *http.Request) {
	goodId := r.FormValue("id")
	goods := &models.Php41Goods{}
	_, err := MasterDB.Id(goodId).Get(goods)
	goodPLus := &models.Php41GoodsPLus{
		Php41Goods: goods,
	}
	goodsIntro := &models.Php41GoodsIntroduce{}
	_, err = MasterDB.Id(goodPLus.GoodsId).Get(goodsIntro)
	fmt.Println(goodsIntro)
	goodPLus.Intro = tools.Stripslashes(goodsIntro.GoodsIntroduce)
	user := &models.Php41Users{}
	_, err = MasterDB.Id(goodPLus.UserId).Get(user)
	checkErr(err)
	data := make(map[string]interface{})
	data["Compose"] = goodPLus
	data["User"] = user
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
	_ = api.ApiRestful(w)

}
func IsOnline(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["is_online"] = 1
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    data,
	}
	_ = api.ApiRestful(w)
}

func AuthorDynamic(w http.ResponseWriter, r *http.Request) {
	userId := r.FormValue("id")
	goods := make([]*models.Php41Goods, 0)
	err := MasterDB.Where("(user_id=?)", userId).Limit(10).Find(&goods)
	checkErr(err)
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    goods,
	}
	err = api.ApiRestful(w)
	checkErr(err)
}

func NiceComment(w http.ResponseWriter, r *http.Request) {
	commentId := r.FormValue("id")
	comments := make([]*models.Php41Ooxx, 0)
	err := MasterDB.Where("(target_id=?)", commentId).Find(&comments)
	fmt.Println(err)
	commentInfos := make([]models.CommentInfo, 10)
	user := make([] *models.Php41Users, 0)
	var temp *models.Php41Ooxx;
	for k, v := range comments {
		err = MasterDB.Where("user_id = ?", v.FromUser).Find(&user)
		fmt.Println(user)
		temp = v;
		commentInfos[k].Php41Ooxx = temp
		commentInfos[k].Author = user[k].Username
		fmt.Println(commentInfos)

	}
	api := initCommon.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    commentInfos,
	}

	err = api.ApiRestful(w)

}

func Demo(w http.ResponseWriter, r *http.Request) {
	users := make([]*models.Php41Users, 0)
	err := MasterDB.Where("(user_id>?)", 0).Find(&users)
	for _, user := range users {
		fmt.Println(user)
	}

	checkErr(err)
	userIdList := make([]int, len(users))
	for key, user := range users {
		userIdList[key] = user.UserId
	}
	goods := make([]*models.Php41Goods, 0)
	err = MasterDB.In("user_id", userIdList).Find(&goods)
	checkErr(err)
}
