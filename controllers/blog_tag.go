package controllers

import (
	"strings"
	"time"

	"pixiublog/models"

	"github.com/astaxie/beego"
)

type BlogTagController struct {
	BaseController
}

func (self *BlogTagController) List() {
	self.display()
}

func (self *BlogTagController) Add() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := models.RoleGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["role_name"] = v.RoleName
		list[k] = row
	}

	self.Data["role"] = list

	self.display()
}

func (self *BlogTagController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogTag, _ := models.BlogTagGetById(id)
	row := make(map[string]interface{})
	row["id"] = blogTag.Id
	row["name"] = blogTag.Name
	row["description"] = blogTag.Description
	self.Data["blogTag"] = row
	self.display()
}

func (self *BlogTagController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	if id == 0 {
		blogTag := new(models.BlogTag)
		blogTag.Name = strings.TrimSpace(self.GetString("name"))
		blogTag.Description = strings.TrimSpace(self.GetString("description"))
		blogTag.CreateTime = time.Now()
		blogTag.UpdateTime = time.Now()

		if _, err := models.BlogTagAdd(blogTag); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	blogTag, _ := models.BlogTagGetById(id)
	//修改
	blogTag.Id = id
	blogTag.Name = strings.TrimSpace(self.GetString("name"))
	blogTag.Description = strings.TrimSpace(self.GetString("description"))
	blogTag.UpdateTime = time.Now()

	if err := blogTag.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogTagController) Del() {
	id, _ := self.GetInt("id")
	blogTag, _ := models.BlogTagGetById(id)
	blogTag.UpdateTime = time.Now()
	blogTag.Status = 1

	if err := blogTag.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
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
	result, count := models.BlogTagGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["name"] = v.Name
		row["description"] = v.Description
		row["createTime"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["updateTime"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
