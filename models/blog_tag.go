package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BlogTag struct {
	Id          int
	Name        string
	CreateTime  time.Time
	UpdateTime  time.Time
	Description string
	Status      int
}

func (t *BlogTag) TableName() string {
	return TableName("biz_tags")
}

func (t *BlogTag) Update(fields ...string) error {
	if t.Name == "" {
		return fmt.Errorf("名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func BlogTagAdd(t *BlogTag) (int64, error) {
	if t.Name == "" {
		return 0, fmt.Errorf("名称不能为空")
	}
	return orm.NewOrm().Insert(t)
}

func BlogTagGetById(id int) (*BlogTag, error) {
	obj := &BlogTag{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func BlogTagDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("biz_tags")).Filter("id", id).Delete()
	return err
}

func BlogTagGetList(page, pageSize int, filters ...interface{}) ([]*BlogTag, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BlogTag, 0)
	query := orm.NewOrm().QueryTable(TableName("biz_tags"))
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
