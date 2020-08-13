package controllers

import (
	"strings"
	"time"

	. "pixiublog/models"
	"pixiublog/utils"
)

type BlogTypeController struct {
	BaseController
}

func (self *BlogTypeController) List() {
	self.display()
}

func (self *BlogTypeController) Add() {
	self.display()
}

func (self *BlogTypeController) Edit() {
	id, _ := self.GetInt("id", 0)
	blogType := &BlogType{}
	blogType.Id = id
	blogType.GetById(blogType)
	self.Data["data"] = utils.StructToJsonThenMap(blogType)
	self.display()
}

func (self *BlogTypeController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	blogType := &BlogType{}

	if id > 0 {
		blogType.Id = id
		blogType.GetById(blogType)
		if blogType == nil {
			self.ajaxMsg("数据不存在", MSG_ERR)
		}
	}

	blogType.Pid, _ = self.GetUint64("pid", 0)
	blogType.Name = strings.TrimSpace(self.GetString("name"))
	blogType.Description = strings.TrimSpace(self.GetString("description"))
	blogType.Sort, _ = self.GetInt("sort", 0)
	blogType.Icon = strings.TrimSpace(self.GetString("icon"))
	blogType.Available, _ = self.GetInt("available", 0)
	blogType.UpdateTime = time.Now()

	if id == 0 {
		blogType.CreateTime = time.Now()
		if _, err := blogType.Add(blogType); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	if err := blogType.Update(blogType); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *BlogTypeController) Del() {
	id, _ := self.GetInt("id")
	blogType := &BlogType{}
	blogType.Id = id
	blogType.GetById(blogType)
	if blogType == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	blogType.UpdateTime = time.Now()
	blogType.Status = 1

	if err := blogType.Update(blogType); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogTypeController) GetList() {
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
	blogType := &BlogType{}
	result := make([]*BlogType, 0)
	count := blogType.GetList(blogType.TableName(), page, self.pageSize, &result, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}
