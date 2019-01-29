package models

import (
	"blog/common/tools"
	"blog/config"
	. "blog/db"
	"blog/redis"
	"encoding/json"
	"fmt"
	"net/http"
)

var BasicUserInfo *UserInfoRedis

type Php41Users struct {
	UserId    int    `json:"user_id" xorm:"pk"`
	Username  string `json:"username"`
	Tel       string `json:"tel"`
	Password  string `json:"password" xorm:"varchar(200)"`
	Taglist   string `json:"taglist" xorm:"varchar(200)"`
	Headimg   string `json:"headimg" xorm:"varchar(200)"`
	Position  string `json:"position" xorm:varchar(200)"`
	Company   string `json:"company" xorm:"varchar(200)"`
	SelfIntro string `json:"self_intro" xorm:"varchar(200)"`
	Homepage  string `json:"homepage" xorm:"varchar(200)"`
	THeadimg  string `json:"t_headimg" xorm:"varchar(200)"`
}

type UserInfoRedis struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"username"`
	Tel       string `json:"tel"`
	PassWord  string `json:"password"`
	TagList   string `json:"taglist"`
	HeadImg   string `json:"headimg"`
	Position  string `json:"position"`
	Company   string `json:"company"`
	SelfIntro string `json:"self_intro"`
	HomePage  string `json:"homepage"`
	THeadImg  string `json:"t_headimg"`
}

func Auth(w http.ResponseWriter, r *http.Request) int {

	cookiePointer, err := r.Cookie("token")
	if err != nil {
		fmt.Println(err)
		return config.ErrAuth
	}
	getValue, err := redis.Get(cookiePointer.Value)
	err = json.Unmarshal([]byte(getValue), &BasicUserInfo)
	if err != nil {
		fmt.Println(err)
		return config.ErrAuth
	}
	return 0
}

func Issue(w http.ResponseWriter, r *http.Request) int {
	cookieValue := tools.GetRandomString(32)
	users := make([]*Php41Users, 0)
	err := MasterDB.Where("(user_id=?)", 30).Find(&users)
	fmt.Println(err)
	if err != nil {
		return config.ErrSourceNotFound
	}
	redisValue, err := json.Marshal(users[0])
	if res := redis.Set(cookieValue, string(redisValue)); res != "OK" {
		fmt.Println(res)
		return config.ErrSetRedisUserInfo
	}

	cookie := http.Cookie{Name: "token", Value: cookieValue, Path: "/", MaxAge: 86400}
	fmt.Println(cookieValue)
	http.SetCookie(w, &cookie)
	return 0
}
