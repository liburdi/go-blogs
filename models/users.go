package models

type Php41Users struct {
	UserId           int    `json:"user_id" xorm:"pk"`
	Username         string `json:"username"`
	Tel              string `json:"tel"`
	Password         string `json:"password" xorm:"varchar(200)"`
	Taglist          string `json:"taglist" xorm:"varchar(200)"`
	Headimg          string `json:"headimg" xorm:"varchar(200)"`
	Position         string `json:"position" xorm:varchar(200)"`
	Company          string `json:"company" xorm:"varchar(200)"`
	SelfIntro string `json:"self_intro" xorm:"varchar(200)"`
	Homepage         string `json:"homepage" xorm:"varchar(200)"`
	THeadimg         string `json:"t_headimg" xorm:"varchar(200)"`
}
