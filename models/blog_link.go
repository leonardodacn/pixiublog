package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type BlogLink struct {
	Base
	Url             string    `json:"url"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Email           string    `json:"email"`
	Qq              string    `json:"qq"`
	Favicon         string    `json:"favicon"`
	Status          int       `json:"-"`
	HomePageDisplay int       `json:"homePageDisplay"`
	Remark          string    `json:"remark"`
	Source          string    `json:"source"`
	CreateTime      time.Time `json:"createTime"`
	UpdateTime      time.Time `json:"updateTime"`
}

func (t *BlogLink) TableName() string {
	return TableName("blog_link")
}

func (t *BlogLink) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(t, fields...); err != nil {
		return err
	}
	return nil
}

func BlogLinkAdd(t *BlogLink) (int64, error) {
	return orm.NewOrm().Insert(t)
}

func BlogLinkGetById(id int) (*BlogLink, error) {
	obj := &BlogLink{}
	obj.Id = id
	err := orm.NewOrm().Read(obj)
	if err != nil {
		return nil, err
	}
	return obj, nil
}

func BlogLinkDelById(id int) error {
	_, err := orm.NewOrm().QueryTable(TableName("blog_link")).Filter("id", id).Delete()
	return err
}

func BlogLinkGetList(page, pageSize int, filters ...interface{}) ([]*BlogLink, int64) {
	offset := (page - 1) * pageSize
	list := make([]*BlogLink, 0)
	query := orm.NewOrm().QueryTable(TableName("blog_link"))
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
