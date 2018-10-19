package models

import "time"

type Php41Goods struct {
	GoodsId        int64     `json:"goods_id" xorm:"pk"`
	GoodsName      string    `json:"goods_name"`
	GoodsNumber    int       `json:"goods_number"`
	GoodsBigLogo   string    `json:"goods_big_logo" xorm:"text"`
	GoodsSmallLogo string    `json:"goods_small_logo" xorm:"char(100)"`
	GoodsIntroduce string    `json:"goods_introduce" xorm:"longtext"`
	IsSale         string    `json:"is_sale" `
	IsRec          string    `json:"is_rec" `
	IsHot          string    `json:"is_hot" `
	IsNew          string    `json:"is_new" `
	AddTime        time.Time `json:"add_time" xorm:"<-"`
	UpdTime        time.Time `json:"add_time" xorm:"<-"`
	IsDel          string    `json:"is_del"`
	Author         string    `json:"author" xorm:"varchar(32)"`
	Pagetitle      string    `json:"pagetitle" xorm:"varchar(100)"`
	Keywords       string    `json:"keywords" xorm:"varchar(300)"`
	Description    string    `json:"description" xorm:"varchar(300)"`
	Cate           string    `json:"cate" xorm:"char(32)"`
	CatId          int       `json:"cat_id" xorm:"int(2)"`
	Yauthor        string    `json:"yauthor" xorm:"varchar(32)"`
	Yurl           string    `json:"yurl" xorm:"varchar(255)"`
	UserId         int       `json:"user_id" xorm:"int(11)"`
	City           string    `json:"city" xorm:"varchar(50)"`
}
