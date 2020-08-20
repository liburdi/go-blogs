package models

type MentalityReview struct {
	Id          int64  `json:"id" xorm:"pk"`
	PicUrl		string  `json:"pic_url" xorm:"varchar(255)"`
	Name        string `json:"name" xorm:"varchar(255)"`
	Desc        string `json:"desc" xorm:"varchar(255)"`
	Type        int8   `json:"type"`
	QuestionIds string `json:"question_ids" xorm:"varchar(255)"`
	CreatedAt   int64  `json:"created_at" xorm:"created"`
	UpdatedAt   int64  `json:"updated_at"  xorm:"updated"`
	IsNice      int8   `json:"is_nice" `
	CreatorId   int64  `json:"creator_id"`
	OperatorId  int64  `json:"operator_id"`
}
