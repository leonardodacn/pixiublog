package controllers

import (
	"strings"
	"time"

	. "pixiublog/models"
	"pixiublog/utils"
)

type BlogCommentController struct {
	BaseController
}

func (self *BlogCommentController) List() {
	self.display()
}

func (self *BlogCommentController) Add() {
	self.display()
}

func (self *BlogCommentController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogComment := &BlogComment{}
	blogComment.Id = id
	blogComment.GetById(blogComment)
	self.Data["data"] = utils.StructToJsonThenMap(blogComment)
	self.display()
}

func (self *BlogCommentController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	blogComment := &BlogComment{}

	if id > 0 {
		blogComment.Id = id
		blogComment.GetById(blogComment)
		if blogComment == nil {
			self.ajaxMsg("数据不存在", MSG_ERR)
		}
	}

	if err := self.ParseForm(blogComment); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	blogComment.UpdateTime = time.Now()

	if id == 0 {
		blogComment.CreateTime = time.Now()
		if _, err := blogComment.Add(blogComment); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	if err := blogComment.Update(blogComment); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogCommentController) Del() {
	id, _ := self.GetInt("id")
	blogComment := &BlogComment{}
	blogComment.Id = id
	blogComment.GetById(blogComment)
	if blogComment == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	blogComment.UpdateTime = time.Now()
	blogComment.Status = 3

	if err := blogComment.Update(blogComment); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogCommentController) GetList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	content := strings.TrimSpace(self.GetString("content"))

	self.pageSize = limit

	filters := make([]interface{}, 0)
	filters = append(filters, "status__lt", 3)
	if content != "" {
		filters = append(filters, "content__icontains", content)
	}
	blogComment := &BlogComment{}
	result := make([]*BlogComment, 0)
	count := blogComment.GetList(blogComment.TableName(), page, self.pageSize, &result, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}
