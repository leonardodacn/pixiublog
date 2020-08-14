package controllers

import (
	"strings"
	"time"

	. "pixiublog/models"
	"pixiublog/utils"
)

type BlogLinkController struct {
	BaseController
}

func (self *BlogLinkController) List() {
	self.display()
}

func (self *BlogLinkController) Add() {
	self.display()
}

func (self *BlogLinkController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogLink := &BlogLink{}
	blogLink.Id = id
	blogLink.GetById(blogLink)
	self.Data["data"] = utils.StructToJsonThenMap(blogLink)
	self.display()
}

func (self *BlogLinkController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	blogLink := &BlogLink{}

	if id > 0 {
		blogLink.Id = id
		blogLink.GetById(blogLink)
		if blogLink == nil {
			self.ajaxMsg("数据不存在", MSG_ERR)
		}
	}

	if err := self.ParseForm(blogLink); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	blogLink.UpdateTime = time.Now()

	if id == 0 {
		blogLink.CreateTime = time.Now()
		if _, err := blogLink.Add(blogLink); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	if err := blogLink.Update(blogLink); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogLinkController) Del() {
	id, _ := self.GetInt("id")
	blogLink := &BlogLink{}
	blogLink.Id = id
	blogLink.GetById(blogLink)
	if blogLink == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	blogLink.UpdateTime = time.Now()
	blogLink.Status = 1

	if err := blogLink.Update(blogLink); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogLinkController) GetList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	name := strings.TrimSpace(self.GetString("name"))

	self.pageSize = limit

	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)
	if name != "" {
		filters = append(filters, "name__icontains", name)
	}
	blogLink := &BlogLink{}
	result := make([]*BlogLink, 0)
	count := blogLink.GetList(blogLink.TableName(), page, self.pageSize, &result, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}
