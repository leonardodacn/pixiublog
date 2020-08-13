package models

import (
	"time"
)

type BlogLink struct {
	Base
	Url             string    `json:"url"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Email           string    `json:"email"`
	Qq              string    `json:"qq"`
	Favicon         string    `json:"favicon"`
	Status          int       `json:"-"`
	HomePageDisplay int       `json:"homePageDisplay"`
	Remark          string    `json:"remark"`
	Source          string    `json:"source"`
	CreateTime      time.Time `json:"createTime"`
	UpdateTime      time.Time `json:"updateTime"`
}

func (t *BlogLink) TableName() string {
	return TableName("blog_link")
}
