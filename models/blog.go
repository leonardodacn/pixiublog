package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type Blog struct {
	Base
	Title       string    `json:"title"       form:"title"`
	UserId      uint64    `json:"userId"     form:"userId"`
	CoverImage  string    `json:"coverImage" form:"coverImage"`
	QrcodePath  string    `json:"qrcodePath" form:"qrcodePath"`
	IsMarkdown  int       `json:"isMarkdown" form:"isMarkdown"`
	Content     string    `json:"content"     form:"content"`
	ContentMd   string    `json:"contentMd"  form:"contentMd"`
	Top         int       `json:"top"         form:"top"`
	TypeId      uint64    `json:"typeId"     form:"typeId"`
	Status      int       `json:"status"      form:"status"`
	Recommended int       `json:"recommended" form:"recommended"`
	Original    int       `json:"original"    form:"original"`
	Description string    `json:"description" form:"description"`
	Keywords    string    `json:"keywords"    form:"keywords"`
	Comment     int       `json:"comment"     form:"comment"`
	CreateTime  time.Time `json:"createTime" form:"createTime"`
	UpdateTime  time.Time `json:"updateTime" form:"updateTime"`
	Code        string    `json:"code"        form:"code"`
	LookCount   int64     `json:"lookCount"  form:"lookCount"`
}

func (t *Blog) TableName() string {
	return TableName("blog")
}

func (t *Blog) UpdateLookCount(id int) {
	_, err := orm.NewOrm().Raw("update "+t.TableName()+" set look_count=look_count + 1 where id = ?", id).Exec()
	if err != nil {
		fmt.Println("DelByBlogId error: {}", err)
	}
}

func (t *Blog) ListByTag(tagID, page, pageSize int, result interface{}, filters ...interface{}) int {
	offset := (page - 1) * pageSize
	_, err := orm.NewOrm().Raw("SELECT  * from blog a left join blog_tag_relation b on(a.id = b.blog_id ) where a.status=1 and b.tag_id = ? order by a.id desc limit ?, ?", tagID, offset, pageSize).QueryRows(result)
	if err != nil {
		fmt.Println("ListByTag error:", err)
	}
	total := 0
	errs := orm.NewOrm().Raw("SELECT  count(1) total from blog a left join blog_tag_relation b on(a.id = b.blog_id ) where b.tag_id = ?", tagID).QueryRow(&total)
	if errs != nil {
		fmt.Println("ListByTag total error:", err)
	}
	return total
}
