package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Category struct {
	Id              int64     `json:"id" gorm:"primary_key"`
	Title           string    `json:"title" sql:"not null;type:varchar(255)"`
	Created         time.Time `json:"created" sql:""`
	Views           int64     `json:"views" sql:"null;type:bigint"`
	TopicTime       time.Time `json:"topictime" sql:""`
	TopicCount      int64     `json:"topiccount" sql:"null;type:bigint"`
	TopicLastUserId int64     `json:"topiclastuserid" sql:"null;type:bigint"`
}
