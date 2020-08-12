package controllers

import (
	"strings"
	"time"

	"fmt"

	"strconv"

	"pixiublog/models"

	"github.com/astaxie/beego"
)

type GroupController struct {
	BaseController
}

func (self *GroupController) List() {
	self.display()
}

func (self *GroupController) Add() {
	self.Data["hideTop"] = true
	self.display()
}
func (self *GroupController) Edit() {
	self.Data["hideTop"] = true
	id, _ := self.GetInt("id", 0)
	group, _ := models.GroupGetById(id)
	row := make(map[string]interface{})
	row["id"] = group.Id
	row["group_name"] = group.GroupName
	row["description"] = group.Description
	self.Data["group"] = row
	self.display()
}

func (self *GroupController) SaveOrUpdate() {
	group := new(models.Group)
	group.GroupName = strings.TrimSpace(self.GetString("group_name"))
	group.Description = strings.TrimSpace(self.GetString("description"))
	group.Status = 1

	group_id, _ := self.GetInt("id")

	fmt.Println(group_id)
	if group_id == 0 {
		//新增
		group.CreateTime = time.Now().Unix()
		group.UpdateTime = time.Now().Unix()
		group.CreateId = self.userId
		group.UpdateId = self.userId
		if _, err := models.GroupAdd(group); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}
	//修改
	group.Id = group_id
	group.UpdateTime = time.Now().Unix()
	group.UpdateId = self.userId
	if err := group.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *GroupController) Del() {

	group_id, _ := self.GetInt("id")
	group, _ := models.GroupGetById(group_id)
	group.Status = 0
	group.Id = group_id
	group.UpdateTime = time.Now().Unix()
	//TODO 如果分组下有任务 不处理
	//filters := make([]interface{}, 0)
	//filters = append(filters, "group_id", group_id)
	//filters = append(filters, "status", 0)
	//_, n := models.TaskServerGetList(1, 1, filters...)
	//if n > 0 {
	//	self.ajaxMsg("分组下有服务器资源，请先处理", MSG_ERR)
	//}
	if err := group.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *GroupController) GetList() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	groupName := strings.TrimSpace(self.GetString("groupName"))
	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)

	if self.userId != 1 {
		groups := strings.Split(self.taskGroups, ",")

		groupsIds := make([]int, 0)
		for _, v := range groups {
			id, _ := strconv.Atoi(v)
			groupsIds = append(groupsIds, id)
		}
		filters = append(filters, "id__in", groupsIds)
	}
	if groupName != "" {
		filters = append(filters, "group_name__contains", groupName)
	}
	result, count := models.GroupGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["group_name"] = v.GroupName
		row["description"] = v.Description
		row["create_time"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["update_time"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}
