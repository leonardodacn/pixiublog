package models

import (
	"time"
)

type BlogTag struct {
	Base
	Name        string    `json:"name"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Description string    `json:"description"`
	Status      int       `-`
}

func (t *BlogTag) TableName() string {
	return TableName("blog_tag")
}
