package models

import (
	"github.com/jinzhu/gorm"
)

type Topic struct {
	ID              int64  `json:"id" gorm:"primary_key"`
	Uid             int64  `json:"uid" sql:"null;type:bigint;default:0"`
	Title           string `json:"title" sql:"not null;type:varchar(255)"`
	Content         string `json:"content" sql:"null;type:longtext"`
	Attachment      string `json:"attachment" sql:"null;type:varchar(255)"`
	Created         string `json:"created" sql:""`
	Updated         string `json:"updated" sql:""`
	Views           int64  `json:"views" sql:"null;type:bigint;default:0",`
	Author          string `json:"author" sql:"null;type:varchar(255);default:'xiechuan'"`
	ReplyTime       string `json:"replytime" sql:""`
	ReplyCount      int64  `json:"replycount" sql:"null;type:bigint;default:0"`
	ReplyLastUserId int64  `json:"replylastuserid" sql:"null;type:bigint;default:0"`
}

func (p *Topic) TableName() string {
	return "topic"
}

func (p *Topic) GetTopic() *gorm.DB {
	return conn.Model(&Topic{})
}
