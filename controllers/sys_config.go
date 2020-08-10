package controllers

import (
	"strings"
	"time"

	"pixiublog/models"

	"github.com/astaxie/beego"
)

type SysConfigController struct {
	BaseController
}

func (self *SysConfigController) List() {
	self.Data["pageTitle"] = "参数配置管理"
	self.display()
}

func (self *SysConfigController) Add() {
	self.Data["pageTitle"] = "新增参数配置"

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

func (self *SysConfigController) Edit() {
	self.Data["pageTitle"] = "编辑参数配置"

	id, _ := self.GetInt("id", 0)
	sysConfig, _ := models.SysConfigGetById(id)
	row := make(map[string]interface{})
	row["id"] = sysConfig.Id
	row["configKey"] = sysConfig.ConfigKey
	row["configValue"] = sysConfig.ConfigValue
	row["configDesc"] = sysConfig.ConfigDesc
	self.Data["sysConfig"] = row
	self.display()
}

func (self *SysConfigController) AjaxSave() {
	id, _ := self.GetInt("id")
	if id == 0 {
		sysConfig := new(models.SysConfig)
		sysConfig.ConfigKey = strings.TrimSpace(self.GetString("configKey"))
		sysConfig.ConfigValue = strings.TrimSpace(self.GetString("configValue"))
		sysConfig.ConfigDesc = strings.TrimSpace(self.GetString("configDesc"))
		sysConfig.CreateTime = time.Now().Unix()
		sysConfig.UpdateTime = time.Now().Unix()

		if _, err := models.SysConfigAdd(sysConfig); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		self.ajaxMsg("", MSG_OK)
	}

	sysConfig, _ := models.SysConfigGetById(id)
	//修改
	sysConfig.Id = id
	sysConfig.ConfigKey = strings.TrimSpace(self.GetString("configKey"))
	sysConfig.ConfigValue = strings.TrimSpace(self.GetString("configValue"))
	sysConfig.ConfigDesc = strings.TrimSpace(self.GetString("configDesc"))
	sysConfig.UpdateTime = time.Now().Unix()

	if err := sysConfig.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("", MSG_OK)
}

func (self *SysConfigController) AjaxDel() {
	id, _ := self.GetInt("id")
	sysConfig, _ := models.SysConfigGetById(id)
	sysConfig.UpdateTime = time.Now().Unix()
	sysConfig.Status = 1

	if err := sysConfig.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *SysConfigController) Table() {
	//列表
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	configKey := strings.TrimSpace(self.GetString("configKey"))

	self.pageSize = limit
	//查询条件
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)
	if configKey != "" {
		filters = append(filters, "config_key__icontains", configKey)
	}
	result, count := models.SysConfigGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["configKey"] = v.ConfigKey
		row["configValue"] = v.ConfigValue
		row["configDesc"] = v.ConfigDesc
		row["createTime"] = beego.Date(time.Unix(v.CreateTime, 0), "Y-m-d H:i:s")
		row["updateTime"] = beego.Date(time.Unix(v.UpdateTime, 0), "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}