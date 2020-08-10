package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type SysConfig struct {
	Id          int
	ConfigKey   string
	ConfigValue string
	CreateTime  int64
	UpdateTime  int64
	ConfigDesc  string
	Status      int
}

func (t *SysConfig) TableName() string {
	return TableName("sys_config")
}

func (t *SysConfig) Update(fields ...string) error {
	if t.ConfigKey == "" {
		return fmt.Errorf("关键字不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func SysConfigAdd(t *SysConfig) (int64, error) {
	if t.ConfigKey == "" {
		return 0, fmt.Errorf("关键字不能为空")
	}
	return orm.NewOrm().Insert(t)
}

func SysConfigGetById(id int) (*SysConfig, error) {
	obj := &SysConfig{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func SysConfigDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("sys_config")).Filter("id", id).Delete()
	return err
}

func SysConfigGetList(page, pageSize int, filters ...interface{}) ([]*SysConfig, int64) {
	offset := (page - 1) * pageSize
	list := make([]*SysConfig, 0)
	query := orm.NewOrm().QueryTable(TableName("sys_config"))
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(&list)
	return list, total
}
