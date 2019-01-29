package controllers

import (
	init_com "blog/common/init"
	. "blog/db"
	"blog/models"
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
	api := init_com.ApiRestful{
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
	ParentChontroller(w, r)
	goodId := r.FormValue("id")
	//goods := make([]*models.Php41Goods, 0)
	goods := &models.Php41Goods{}
	_, err := MasterDB.Id(goodId).Get(goods)
	goodPLus := &models.Php41GoodsPLus{
		Php41Goods: goods,
	}
	goodsIntro := &models.Php41GoodsIntroduce{}
	_, err = MasterDB.Id(goodPLus.GoodsId).Get(goodsIntro)
	goodPLus.Intro = goodsIntro.GoodsIntroduce
	checkErr(err)
	api := init_com.ApiRestful{
		Code:    200,
		Message: "Success",
		Data:    goodPLus,
	}
	_ = api.ApiRestful(w)
	//checkErr(err)

}
func TestConn(w http.ResponseWriter, r *http.Request) {
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
	//for _, good := range goods {
	//	fmt.Println(good.UserId)
	//}
}
