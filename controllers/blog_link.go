package controllers

import (
	"strings"
	"time"

	"pixiublog/models"
	"pixiublog/utils"
)

type BlogLinkController struct {
	BaseController
}

func (self *BlogLinkController) List() {
	self.display()
}

func (self *BlogLinkController) Add() {
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

func (self *BlogLinkController) Edit() {
	id, _ := self.GetInt("id", 0)
	//blogLink, _ := models.BlogLinkGetById(id)
	blogLink := &models.BlogLink{}
	b, _ := blogLink.GetById(id)
	self.Data["blogLink"] = utils.StructToJsonThenMap(*b)
	self.display()
}

func (self *BlogLinkController) AjaxSave() {
	id, _ := self.GetInt("id")
	blogLink := new(models.BlogLink)

	if id > 0 {
		blogLink, _ = models.BlogLinkGetById(id)
	}

	blogLink.Url = self.GetString("url")
	blogLink.Name = self.GetString("name")
	blogLink.Description = self.GetString("description")
	blogLink.Email = self.GetString("email")
	blogLink.Qq = self.GetString("qq")
	blogLink.Favicon = self.GetString("favicon")
	blogLink.HomePageDisplay = self.GetAlwaysInt("homePageDisplay")
	blogLink.Remark = self.GetString("remark")
	blogLink.Source = self.GetString("source")

	if id == 0 {
		blogLink.CreateTime = time.Now()
		blogLink.UpdateTime = time.Now()

		if _, err := models.BlogLinkAdd(blogLink); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	blogLink.UpdateTime = time.Now()

	if err := blogLink.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogLinkController) AjaxDel() {
	id, _ := self.GetInt("id")
	blogLink, _ := models.BlogLinkGetById(id)
	blogLink.UpdateTime = time.Now()
	blogLink.Status = 1

	if err := blogLink.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogLinkController) Table() {
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
	result, count := models.BlogLinkGetList(page, self.pageSize, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}
