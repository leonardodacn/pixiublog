package models

import (
	"time"
)

type BlogLink struct {
	Base
	Name            string    `json:"name"            form:"name"`
	Url             string    `json:"url"             form:"url"`
	Description     string    `json:"description"     form:"description"`
	Email           string    `json:"email"           form:"email"`
	Qq              string    `json:"qq"              form:"qq"`
	Favicon         string    `json:"favicon"         form:"favicon"`
	Status          int       `json:"-"               form:"-"`
	HomePageDisplay int       `json:"homePageDisplay" form:"homePageDisplay"`
	Remark          string    `json:"remark"          form:"remark"`
	Source          string    `json:"source"          form:"source"`
	CreateTime      time.Time `json:"createTime"      form:"createTime"`
	UpdateTime      time.Time `json:"updateTime"      form:"updateTime"`
}

func (t *BlogLink) TableName() string {
	return TableName("blog_link")
}
