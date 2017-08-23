package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Topic struct {
	ID              int64     `json:"id" gorm:"primary_key"`
	Uid             int64     `json:"uid" sql:"null;type:bigint"`
	Title           string    `json:"title" sql:"not null;type:varchar(255)"`
	Content         string    `json:"content" sql:"not null;type:varchar(255)"`
	Attachment      string    `json:"attachment" sql:"not null;type:varchar(255)"`
	Created         time.Time `json:"created" sql:""`
	Updated         time.Time `json:"updated" sql:""`
	Views           int64     `json:"views" sql:"null;type:bigint"`
	Author          string    `json:"author" sql:"not null;type:varchar(255)"`
	ReplyTime       time.Time `json:"replytime" sql:""`
	ReplyCount      int64     `json:"replycount" sql:"null;type:bigint"`
	ReplyLastUserId int64     `json:"replylastuserid" sql:"null;type:bigint"`
}

func (p *Topic) TableName() string {
	return "topic"
}

func (p *Topic) GetTopic() *gorm.DB {
	return conn.Model(&Topic{})
}
