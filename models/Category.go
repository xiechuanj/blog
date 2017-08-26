package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	Id              int64  `json:"id" gorm:"primary_key"`
	Title           string `json:"title" sql:"not null;type:varchar(255)"`
	Created         string `json:"created" sql:""`
	Views           int64  `json:"views" sql:"null;type:bigint;default:0"`
	TopicTime       string `json:"topictime" sql:""`
	TopicCount      int64  `json:"topiccount" sql:"null;type:bigint;default:0"`
	TopicLastUserId int64  `json:"topiclastuserid" sql:"null;type:bigint;default:0"`
}

func (p *Category) TableName() string {
	return "category"
}

func (p *Category) GetCategory() *gorm.DB {
	return conn.Model(&Category{})
}
