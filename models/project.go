package models


type Php41Project struct {
	Id          int       `json:"id" xorm:"pk"`
	Title       string    `json:"title" xorm:"varchar(255) notnull"`
	Desc         string    `json:"desc" xorm:"varchar(255) default('')""`
	OriginUrl         string    `json:"origin_url" xorm:"varchar(255) default('')""`
	IsDel  int       `json:"is_del" xorm:"tinyint notnull  default(0)"`
	CreatorId  int       `json:"creator_id" xorm:"int notnull  default(0)"`
	CreateTime  int `json:"create_time" xorm:"created"`
	UpdateTime  int `json:"Update_time" xorm:"updated"`
}
