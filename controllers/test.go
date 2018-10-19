package controllers

import (
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
func TestConn(w http.ResponseWriter, r *http.Request) {
	users := make([]*models.Php41Users, 0)
	err := MasterDB.Where("(user_id=?)", 1).Find(&users)
	fmt.Println(err)
	fmt.Println(users)
	userIdList := make([]int, len(users))
	fmt.Println(userIdList)
	for key, user := range users {
		fmt.Println(user.UserId)
		userIdList[key] = user.UserId
		fmt.Println(userIdList[key])
	}
	goods := make([]*models.Php41Goods, 0)
	err = MasterDB.In("user_id", userIdList).Find(&goods)
	fmt.Println(err)
	for _, good := range goods {
		fmt.Println(good)
	}

}
