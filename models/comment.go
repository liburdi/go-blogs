package models

import "time"

type Php41Ooxx struct {
	Id          int       `json:"goods_id" xorm:"pk"`
	TargetId    int       `json:"target_id"`
	TargetType  int       `json:"target_type"`
	Msg         string    `json:"msg" xorm:"varchar(255)"`
	LikeCount   int       `json:"like_count"`
	UnlikeCount int       `json:"unlike_count" `
	AddTime     time.Time `json:"add_time" xorm:"<-"`
	FromUser    int       `json:"from_user"`
	ToUser      int       `json:"to_user"`
	ParentId    int       `json:"parent_id" xorm:"varchar(32)"`
}

type CommentInfo struct {
	*Php41Ooxx
	Author string
}
