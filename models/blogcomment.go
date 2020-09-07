package models

import (
	"time"
)

type BlogComment struct {
	Base
	Sid              int64     `json:"sid"                form:"sid"`
	UserId           uint64    `json:"userId"             form:"userId"`
	Pid              uint64    `json:"pid"                form:"pid"`
	Qq               string    `json:"qq"                 form:"qq"`
	Nickname         string    `json:"nickname"           form:"nickname"`
	Avatar           string    `json:"avatar"             form:"avatar"`
	Email            string    `json:"email"              form:"email"`
	Url              string    `json:"url"                form:"url"`
	Status           int       `json:"status"             form:"status"`
	Ip               string    `json:"ip"                 form:"ip"`
	Lng              string    `json:"lng"                form:"lng"`
	Lat              string    `json:"lat"                form:"lat"`
	Address          string    `json:"address"            form:"address"`
	Os               string    `json:"os"                 form:"os"`
	OsShortName      string    `json:"osShortName"        form:"osShortTame"`
	Browser          string    `json:"browser"            form:"browser"`
	BrowserShortName string    `json:"browserShortName"   form:"browserShortName"`
	Content          string    `json:"content"            form:"content"`
	Remark           string    `json:"remark"             form:"remark"`
	Support          uint      `json:"support"            form:"support"`
	Oppose           uint      `json:"oppose"             form:"oppose"`
	CreateTime       time.Time `json:"createTime"         form:"createTime"`
	UpdateTime       time.Time `json:"updateTime"         form:"updateTime"`
}

func (t *BlogComment) TableName() string {
	return "blog_comment"
}
