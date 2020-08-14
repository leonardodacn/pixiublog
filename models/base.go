package models

import (
	"github.com/astaxie/beego/orm"
)

type Base struct {
	Id int `json:"id" form:"id"`
}

func (t *Base) Update(obj interface{}, fields ...string) error {
	if _, err := orm.NewOrm().Update(obj, fields...); err != nil {
		return err
	}
	return nil
}

func (t *Base) Add(obj interface{}) (int64, error) {
	return orm.NewOrm().Insert(obj)
}

func (t *Base) GetById(obj interface{}) {
	orm.NewOrm().Read(obj)
}

func (t *Base) DelById(obj interface{}) error {
	_, err := orm.NewOrm().Delete(obj)
	return err
}

func (t *Base) GetList(tableName string, page, pageSize int, result interface{}, filters ...interface{}) int64 {
	offset := (page - 1) * pageSize
	query := orm.NewOrm().QueryTable(tableName)
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("-id").Limit(pageSize, offset).All(result)
	return total
}
