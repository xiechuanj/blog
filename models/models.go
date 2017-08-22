package models

import (
	"time"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time
	Views           int64
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string
	Attachment      string
	Created         time.Time
	Updated         time.Time
	Views           int64
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}
