package models

import (
	"encoding/json"
	"fmt"
	"github.com/liburdi/go-blogs/config"
	"github.com/liburdi/go-blogs/redis"
	"net/http"
)

var BasicUserInfo *UserInfoRedis

type Users struct {
	UserId    int    `json:"user_id" xorm:"pk"`
	Username  string `json:"username"`
	Tel       string `json:"tel"`
	Password  string `json:"password" xorm:"varchar(200)"`
	TagList   string `json:"tag_list" xorm:"varchar(200)"`
	HeadImg   string `json:"head_img" xorm:"varchar(200)"`
	Position  string `json:"position" xorm:"varchar(200)"`
	Company   string `json:"company" xorm:"varchar(200)"`
	SelfIntro string `json:"self_intro" xorm:"varchar(200)"`
	Homepage  string `json:"homepage" xorm:"varchar(200)"`
	THeadImg  string `json:"t_head_img" xorm:"varchar(200)"`
}

type UserInfoRedis struct {
	UserId    int    `json:"user_id"`
	Username  string `json:"username"`
	Tel       string `json:"tel"`
	PassWord  string `json:"password"`
	TagList   string `json:"tag_list"`
	HeadImg   string `json:"head_img"`
	Position  string `json:"position"`
	Company   string `json:"company"`
	SelfIntro string `json:"self_intro"`
	HomePage  string `json:"homepage"`
	THeadImg  string `json:"t_head_img"`
}

func Auth(w http.ResponseWriter, r *http.Request) error {
	if r == nil {
		return fmt.Errorf(config.ErrorMessage[config.ErrAuth])
	}
	cookiePointer, err := r.Cookie("token")
	if err != nil {
		return fmt.Errorf(config.ErrorMessage[config.ErrAuth])
	}
	getValue, err := redis.Get(cookiePointer.Value)
	err = json.Unmarshal([]byte(getValue), &BasicUserInfo)
	if err != nil {
		return fmt.Errorf(config.ErrorMessage[config.ErrAuth])
	}
	return nil
}
