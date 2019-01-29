package models

import "time"

type Php41Goods struct {
	GoodsId        int64     `json:"goods_id" xorm:"pk"`
	GoodsName      string    `json:"goods_name"`
	GoodsNumber    int       `json:"goods_number"`
	GoodsSmallLogo string    `json:"goods_small_logo" xorm:"char(100)"`
	IsHot          string    `json:"is_hot"  xorm:"enum('热销','不热销')"`
	IsNew          string    `json:"is_new" `
	AddTime        time.Time `json:"add_time" xorm:"<-"`
	UpdTime        time.Time `json:"upd_time" xorm:"<-"`
	IsDel          string    `json:"is_del"`
	Author         string    `json:"author" xorm:"varchar(32)"`
	Keywords       string    `json:"keywords" xorm:"varchar(300)"`
	Description    string    `json:"description" xorm:"varchar(300)"`
	Cate           string    `json:"cate" xorm:"char(32)"`
	CatId          int       `json:"cat_id" xorm:"int(2)"`
	OriginAuthor   string    `json:"origin_author" xorm:"varchar(32)"`
	OriginUrl      string    `json:"origin_url" xorm:"varchar(255)"`
	UserId         int       `json:"user_id" xorm:"int(11)"`
	City           string    `json:"city" xorm:"varchar(50)"`
}
type Php41GoodsPLus struct {
	*Php41Goods
	Intro string
}
type Php41GoodsIntroduce struct {
	GoodsId        int    `json:"goods_id" xorm:"pk"`
	GoodsIntroduce string `json:"goods_introduce" xorm:"text"`
}
