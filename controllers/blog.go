package controllers

import (
	"strconv"
	"strings"
	"time"

	. "pixiublog/models"
	"pixiublog/utils"
)

type BlogController struct {
	BaseController
}

func (self *BlogController) List() {
	self.display()
}

func (self *BlogController) Add() {
	setData(self)
	self.display()
}

func (self *BlogController) Edit() {
	id, _ := self.GetInt("id", 0)
	blog := &Blog{}
	blog.Id = id
	blog.GetById(blog)
	self.Data["data"] = utils.StructToJsonThenMap(blog)

	filters := make([]interface{}, 0)
	filters = append(filters, "blog_id", id)
	blogTagRelation := &BlogTagRelation{}
	blogTagRelations := make([]*BlogTagRelation, 0)
	blogTagRelation.GetAll(blogTagRelation.TableName(), &blogTagRelations, "", filters...)
	tags := ""
	if blogTagRelations != nil && len(blogTagRelations) > 0 {
		for i, v := range blogTagRelations {
			if i > 0 {
				tags += ","
			}
			tags += strconv.Itoa(v.TagId)
		}
	}
	self.Data["selectBlogTags"] = tags

	setData(self)

	self.display()
}

func setData(self *BlogController) {
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 0)

	blogType := &BlogType{}
	bolgTypes := make([]*BlogType, 0)
	blogType.GetAll(blogType.TableName(), &bolgTypes, "", filters...)
	self.Data["blogTypes"] = bolgTypes

	blogTag := &BlogTag{}
	blogTags := make([]*BlogTag, 0)
	blogTag.GetAll(blogTag.TableName(), &blogTags, "", filters...)
	self.Data["blogTags"] = blogTags
}

func (self *BlogController) SaveOrUpdate() {
	id, _ := self.GetInt("id")
	blog := &Blog{}

	if id > 0 {
		blog.Id = id
		blog.GetById(blog)
		if blog == nil {
			self.ajaxMsg("数据不存在", MSG_ERR)
		}
	}

	if err := self.ParseForm(blog); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}

	if utils.IsEmpty(blog.Code) {
		blog.Code = utils.GetRandomString(6)
	}
	blog.UpdateTime = time.Now()

	//标签处理
	tags := self.GetString("tags")

	if id == 0 {
		blog.CreateTime = time.Now()
		if blogID, err := blog.Add(blog); err != nil {
			self.ajaxMsg(err.Error(), MSG_ERR)
		} else {
			dealTags(tags, int(blogID))
		}
		clearBlogs()
		self.ajaxMsg("", MSG_OK)
	}

	//修改
	if err := blog.Update(blog); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearBlogs()
	dealTags(tags, blog.Id)
	self.ajaxMsg("", MSG_OK)
}

func dealTags(tags string, blogID int) {
	if !utils.IsEmpty(tags) {
		tagArray := strings.Split(tags, ",")
		blogTagRelation := &BlogTagRelation{}
		blogTagRelation.DelByBlogId(blogID)
		for _, v := range tagArray {
			if tagID, err := strconv.Atoi(v); err == nil {
				blogTagRelation := &BlogTagRelation{}
				blogTagRelation.BlogId = blogID
				blogTagRelation.TagId = tagID
				blogTagRelation.CreateTime = time.Now()
				blogTagRelation.UpdateTime = time.Now()
				blogTagRelation.Add(blogTagRelation)
			}
		}
		clearTags()
	}
}

func (self *BlogController) Del() {
	id, _ := self.GetInt("id")
	blog := &Blog{}
	blog.Id = id
	blog.GetById(blog)
	if blog == nil {
		self.ajaxMsg("数据不存在", MSG_ERR)
	}
	blog.UpdateTime = time.Now()
	blog.Status = 3

	if err := blog.Update(blog); err != nil {
		self.ajaxMsg(err.Error(), MSG_ERR)
	}
	clearBlogs()
	self.ajaxMsg("操作成功", MSG_OK)
}

func (self *BlogController) GetList() {
	page, err := self.GetInt("page")
	if err != nil {
		page = 1
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 30
	}

	content := strings.TrimSpace(self.GetString("content"))

	self.pageSize = limit

	filters := make([]interface{}, 0)
	if status, err := self.GetInt("status"); err == nil && status >= 0 {
		filters = append(filters, "status", status)
	} else {
		filters = append(filters, "status__lt", 3)
	}
	if content != "" {
		filters = append(filters, "content__icontains", content)
	}
	blog := &Blog{}
	result := make([]*Blog, 0)
	count := blog.GetList(blog.TableName(), page, self.pageSize, &result, filters...)
	self.ajaxList("成功", MSG_OK, count, result)
}

var gBlogs = make([]*Blog, 0)

func getAllBlogs() []*Blog {
	if len(gBlogs) > 0 {
		return gBlogs
	}
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	blog := &Blog{}
	blog.GetAll(blog.TableName(), &gBlogs, "", filters...)
	return gBlogs
}

func clearBlogs() {
	gBlogs = gBlogs[0:0]
}
