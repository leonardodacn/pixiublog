package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type BlogTagRelation struct {
	Base
	TagId      int
	BlogId     int
	CreateTime time.Time
	UpdateTime time.Time
}

func (t *BlogTagRelation) TableName() string {
	return TableName("blog_tag_relation")
}

func (t *BlogTagRelation) DelByBlogId(blogID int) {
	_, err := orm.NewOrm().Raw("delete from "+t.TableName()+" where blog_id = ?", blogID).Exec()
	if err != nil {
		fmt.Println("DelByBlogId error: {}", err)
	}
}
