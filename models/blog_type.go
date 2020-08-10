package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BlogType struct {
	Id          int
	Pid         uint64
	Name        string
	Description string
	Sort        int
	Icon        string
	Available   int
	CreateTime  time.Time
	UpdateTime  time.Time
	Status      int
}

func (t *BlogType) TableName() string {
	return TableName("biz_type")
}

func (t *BlogType) Update(fields ...string) error {
	if t.Name == "" {
		return fmt.Errorf("名称不能为空")
	}
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func BlogTypeAdd(t *BlogType) (int64, error) {
	if t.Name == "" {
		return 0, fmt.Errorf("名称不能为空")
	}
	return orm.NewOrm().Insert(t)
}

func BlogTypeGetById(id int) (*BlogType, error) {
	obj := &BlogType{
		Id: id,
	}
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func BlogTypeDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("biz_type")).Filter("id", id).Delete()
	return err
}

func BlogTypeGetList(page, pageSize int, filters ...interface{}) ([]*BlogType, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BlogType, 0)
	query := orm.NewOrm().QueryTable(TableName("biz_type"))
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
