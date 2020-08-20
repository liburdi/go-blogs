package models

type MentalityQuestion struct {
	Id         int64     `json:"id" xorm:"pk"`
	Name       string    `json:"name" xorm:"varchar(255)"`
	Desc       string    `json:"desc" xorm:"varchar(255)"`
	IsMore     int8      `json:"is_more" `
	CreatorId  int64     `json:"creator_id"`
	OperatorId int64     `json:"operator_id"`
}
