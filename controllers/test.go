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
