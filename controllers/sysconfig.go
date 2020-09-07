package controllers

import (
	"strings"
	"time"

	. "pixiublog/models"

	"github.com/astaxie/beego"
)

type SysConfigController struct {
	BaseController
}

func (self *SysConfigController) List() {
	self.display()
}

func (self *SysConfigController) Add() {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	result, _ := RoleGetList(1, 1000, filters...)
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
	id, _ := self.GetInt("id", 0)
	sysConfig, _ := SysConfigGetById(id)
	row := make(map[string]interface{})
	row["id"] = sysConfig.Id
	row["configKey"] = sysConfig.ConfigKey
	row["configValue"] = sysConfig.ConfigValue
	row["configDesc"] = sysConfig.ConfigDesc
	self.Data["sysConfig"] = row
	self.display()
}

func (self *SysConfigController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	if id == 0 {
		sysConfig := new(SysConfig)
		sysConfig.ConfigKey = strings.TrimSpace(self.GetString("configKey"))
		sysConfig.ConfigValue = strings.TrimSpace(self.GetString("configValue"))
		sysConfig.ConfigDesc = strings.TrimSpace(self.GetString("configDesc"))
		sysConfig.CreateTime = time.Now()
		sysConfig.UpdateTime = time.Now()

		if _, err := SysConfigAdd(sysConfig); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		}
		clearSysConfigs()
		self.ajaxMsg("", MSG_OK)
	}

	sysConfig, _ := SysConfigGetById(id)
	//修改
	sysConfig.Id = id
	sysConfig.ConfigKey = strings.TrimSpace(self.GetString("configKey"))
	sysConfig.ConfigValue = strings.TrimSpace(self.GetString("configValue"))
	sysConfig.ConfigDesc = strings.TrimSpace(self.GetString("configDesc"))
	sysConfig.UpdateTime = time.Now()

	if err := sysConfig.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearSysConfigs()
	self.ajaxMsg("", MSG_OK)
}

func (self *SysConfigController) Del() {
	id, _ := self.GetInt("id")
	sysConfig, _ := SysConfigGetById(id)
	sysConfig.UpdateTime = time.Now()
	sysConfig.Status = 1

	if err := sysConfig.Update(); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearSysConfigs()
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *SysConfigController) GetList() {
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
	result, count := SysConfigGetList(page, self.pageSize, filters...)
	list := make([]map[string]interface{}, len(result))
	for k, v := range result {
		row := make(map[string]interface{})
		row["id"] = v.Id
		row["configKey"] = v.ConfigKey
		row["configValue"] = v.ConfigValue
		row["configDesc"] = v.ConfigDesc
		row["createTime"] = beego.Date(v.CreateTime, "Y-m-d H:i:s")
		row["updateTime"] = beego.Date(v.UpdateTime, "Y-m-d H:i:s")
		list[k] = row
	}
	self.ajaxList("成功", MSG_OK, count, list)
}

var gSysConfigs = make(map[string]string)

func getAllSysConfig() map[string]string {
	if len(gSysConfigs) > 0 {
		return gSysConfigs
	}
	sysConfig := &SysConfig{}
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)
	configs := make([]*SysConfig, 0)
	sysConfig.GetAll(sysConfig.TableName(), &configs, "", filters...)
	if len(configs) > 0 {
		for _, v := range configs {
			gSysConfigs[v.ConfigKey] = v.ConfigValue
		}
	}
	return gSysConfigs
}

func clearSysConfigs() {
	gSysConfigs = make(map[string]string)
}
