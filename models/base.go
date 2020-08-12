package models

import (
	"github.com/astaxie/beego/orm"
)

type Base struct {
	Id int
}

func (t *Base) TableName() string {
	return TableName("base")
}

func (t *Base) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Base) Add() (int64, error) {
	return orm.NewOrm().Insert(t)
}

func (t *Base) GetById(id int) (*Base, error) {
	obj := &Base{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func (t *Base) DelById(id int) error {
	_, err := orm.NewOrm().QueryTable(t.TableName()).Filter("id", id).Delete()
	return err
}

func (t *Base) GetList(page, pageSize int, filters ...interface{}) ([]*Base, int64) {
	offset := (page - 1) * pageSize
	list := make([]*Base, 0)
	query := orm.NewOrm().QueryTable(t.TableName())
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
