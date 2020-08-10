package controllers

import (
	"strings"
	"time"

	"pixiublog/models"

	"github.com/astaxie/beego"
)

type BlogTypeController struct {
	BaseController
}

func (self *BlogTypeController) List() {
	self.display()
}

func (self *BlogTypeController) Add() {
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

func (self *BlogTypeController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogType, _ := models.BlogTypeGetById(id)
	row := make(map[string]interface{})
	row["id"] = blogType.Id
	row["pid"] = blogType.Pid
	row["name"] = blogType.Name
	row["description"] = blogType.Description
	row["sort"] = blogType.Sort
	row["icon"] = blogType.Icon
	row["available"] = blogType.Available
	row["createTime"] = blogType.CreateTime
	row["updateTime"] = blogType.UpdateTime
	self.Data["blogType"] = row
	self.display()
}

func (self *BlogTypeController) AjaxSave() {
	id, _ := self.GetInt("id")
	if id == 0 {
		blogType := new(models.BlogType)
		blogType.Pid, _ = self.GetUint64("pid", 0)
		blogType.Name = strings.TrimSpace(self.GetString("name"))
		blogType.Description = strings.TrimSpace(self.GetString("description"))
		blogType.Sort, _ = self.GetInt("sort", 0)
		blogType.Icon = strings.TrimSpace(self.GetString("icon"))
		blogType.Available, _ = self.GetInt("available", 0)
		blogType.CreateTime = time.Now()
		blogType.UpdateTime = time.Now()

		if _, err := models.BlogTypeAdd(blogType); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	blogType, _ := models.BlogTypeGetById(id)
	//修改
	blogType.Id = id
	blogType.Pid, _ = self.GetUint64("pid", 0)
	blogType.Name = strings.TrimSpace(self.GetString("name"))
	blogType.Description = strings.TrimSpace(self.GetString("description"))
	blogType.Sort, _ = self.GetInt("sort", 0)
	blogType.Icon = strings.TrimSpace(self.GetString("icon"))
	blogType.Available, _ = self.GetInt("available", 0)
	blogType.UpdateTime = time.Now()

	if err := blogType.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogTypeController) AjaxDel() {
	id, _ := self.GetInt("id")
	blogType, _ := models.BlogTypeGetById(id)
	blogType.UpdateTime = time.Now()
	blogType.Status = 1

	if err := blogType.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogTypeController) Table() {
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
	result, count := models.BlogTypeGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["pid"] = v.Pid
		row["name"] = v.Name
		row["description"] = v.Description
		row["sort"] = v.Sort
		row["icon"] = v.Icon
		row["available"] = v.Available
		row["createTime"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		row["updateTime"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
