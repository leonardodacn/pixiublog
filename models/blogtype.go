package models

import (
	"time"
)

type BlogType struct {
	Base
	Pid         uint64    `json:"pid"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Sort        int       `json:"sort"`
	Icon        string    `json:"icon"`
	Available   int       `json:"available"`
	CreateTime  time.Time `json:"createTime"`
	UpdateTime  time.Time `json:"updateTime"`
	Status      int       `-`
}

func (t *BlogType) TableName() string {
	return TableName("blog_type")
}
