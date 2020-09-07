package controllers

import (
	"strings"
	"time"

	. "pixiublog/models"
)

type BlogTagController struct {
	BaseController
}

func (self *BlogTagController) List() {
	self.display()
}

func (self *BlogTagController) Add() {
	self.display()
}

func (self *BlogTagController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogTag := &BlogTag{}
	blogTag.Id = id
	blogTag.GetById(blogTag)
	if blogTag == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	row := make(map[string]interface{})
	row["id"] = blogTag.Id
	row["name"] = blogTag.Name
	row["description"] = blogTag.Description
	self.Data["data"] = row
	self.display()
}

func (self *BlogTagController) SaveOrUpdate() {
	id, _ := self.GetInt("id")

	blogTag := &BlogTag{}
	if id > 0 {
		blogTag.Id = id
		blogTag.GetById(blogTag)
		if blogTag == nil {
			self.ajaxMsg("数据不存在", MSG_ERR)
		}
	}

	blogTag.Name = strings.TrimSpace(self.GetString("name"))
	blogTag.Description = strings.TrimSpace(self.GetString("description"))
	blogTag.UpdateTime = time.Now()

	if id == 0 {
		blogTag.CreateTime = time.Now()
		if _, err := blogTag.Add(blogTag); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		clearTags()
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	if err := blogTag.Update(blogTag); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearTags()
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogTagController) Del() {
	id, _ := self.GetInt("id")
	blogTag := &BlogTag{}
	blogTag.Id = id
	blogTag.GetById(blogTag)
	if blogTag == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	blogTag.UpdateTime = time.Now()
	blogTag.Status = 1

	if err := blogTag.Update(blogTag); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearTags()
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogTagController) GetList() {
	//列表
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
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)
	if name != "" {
		filters = append(filters, "name__icontains", name)
	}

	blogTag := &BlogTag{}
	result := make([]*BlogTag, 0)
	count := blogTag.GetList(blogTag.TableName(), page, self.pageSize, &result, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}

var gBlogTags = make([]*BlogTag, 0)

func getAllTags() []*BlogTag {
	if len(gBlogTags) > 0 {
		return gBlogTags
	}
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)
	blogTag := &BlogTag{}
	blogTag.GetAll(blogTag.TableName(), &gBlogTags, "", filters...)
	return gBlogTags
}

func clearTags() {
	gBlogTags = gBlogTags[0:0]
}

func getTagById(id int) *BlogTag {
	blogTags := getAllTags()
	if len(blogTags) > 0 {
		for _, v := range blogTags {
			if v.Id == id {
				return v
			}
		}
	}
	return nil
}
