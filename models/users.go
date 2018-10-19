package models

type Php41Users struct {
	UserId           int    `json:"user_id" xorm:"pk"`
	Username         string `json:"username"`
	Tel              string `json:"tel"`
	Password         string `json:"password" xorm:"varchar(200)"`
	Taglist          string `json:"taglist" xorm:"varchar(200)"`
	Headimg          string `json:"headimg" xorm:"varchar(200)"`
	Position         string `json:"position" xorm:varchar(200)"`
	Company          string `json:"position" xorm:"varchar(200)"`
	Selfintroduction string `json:"selfintroduction" xorm:"varchar(200)"`
	Homepage         string `json:"position" xorm:"varchar(200)"`
	THeadimg         string `json:"position" xorm:"varchar(200)"`
}
