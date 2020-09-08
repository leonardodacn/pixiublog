package controllers

import (
	"html/template"
	. "pixiublog/models"
	"pixiublog/utils"
	"strings"
)

type WebController struct {
	BaseController
}

func (self *WebController) Type() {
	typeID, err := self.GetInt(":type")
	if err != nil {
		typeID = -1
	}
	home(self, typeID, -1, -1)
	self.Data["url"] = "type"
	self.displayWeb("web/index")
}

func (self *WebController) Tag() {
	tagID, err := self.GetInt(":tag")
	if err != nil {
		tagID = -1
	}
	home(self, -1, tagID, -1)
	self.Data["url"] = "tag"
	self.displayWeb("web/index")
}

func (self *WebController) Recommended() {
	home(self, -1, -1, 1)
	self.Data["url"] = "recommended"
	self.displayWeb("web/index")
}

func (self *WebController) Index() {
	home(self, -1, -1, -1)
	self.Data["url"] = "index"
	self.displayWeb()
}

func home(self *WebController, typeID, tagID, recommended int) {
	page, err := self.GetInt(":page")
	if err != nil {
		page, err = self.GetInt("pageNumber")
		if err != nil {
			page = 1
		}
	}
	limit, err := self.GetInt("limit")
	if err != nil {
		limit = 10
	}

	self.pageSize = limit

	filters := make([]interface{}, 0)

	total := 0
	if tagID > 0 {
		blog := &Blog{}
		blogs := make([]*Blog, 0)
		total = blog.ListByTag(tagID, page, self.pageSize, &blogs, filters...)
		self.Data["blogs"] = blogs
	} else {
		filters = append(filters, "status", 1)
		if typeID > 0 {
			filters = append(filters, "type_id", typeID)
		}

		if recommended > 0 {
			filters = append(filters, "recommended", 1)
		}
		content := strings.TrimSpace(self.GetString("keywords"))
		if len(content) == 0 {
			content = strings.TrimSpace(self.GetString("kw"))
		}
		if len(content) > 0 {
			self.Data["keywords"] = content
			self.Data["isSearch"] = true
			filters = append(filters, "content__icontains", content)
		}

		blog := &Blog{}
		blogs := make([]*Blog, 0)
		total = int(blog.GetList(blog.TableName(), page, self.pageSize, &blogs, filters...))
		self.Data["blogs"] = blogs
	}

	pageInfo := &utils.Page{}
	pageInfo.Init(limit, page, int(total))
	self.Data["page"] = pageInfo

	setHomeData(self)
}

func (self *WebController) About() {
	setHomeData(self)
	self.displayWeb()
}

func (self *WebController) GetByCode() {
	setHomeData(self)

	code := self.GetString(":code")
	blog := &Blog{}
	blogFileters := make([]interface{}, 0)
	blogFileters = append(blogFileters, "code", code)
	blogFileters = append(blogFileters, "status__gt", 0)
	blogFileters = append(blogFileters, "status__lt", 3)
	blog.GetOne(blog.TableName(), blog, blogFileters...)

	if blog.Id == 0 {
		self.Abort("404")
	}

	self.Data["blog"] = blog
	self.Data["content"] = template.HTML(blog.Content)
	self.Data["isBlog"] = true

	filters := make([]interface{}, 0)
	filters = append(filters, "blog_id", blog.Id)
	blogTagRelation := &BlogTagRelation{}
	blogTagRelations := make([]*BlogTagRelation, 0)
	blogTagRelation.GetAll(blogTagRelation.TableName(), &blogTagRelations, "", filters...)
	if len(blogTagRelations) > 0 {
		selectBlogTags := make([]*BlogTag, 0)
		for _, v := range blogTagRelations {
			if d := getTagById(v.TagId); d != nil {
				selectBlogTags = append(selectBlogTags, d)
			}

		}
		self.Data["selectBlogTags"] = selectBlogTags
	}

	if blogType := getTypeById(int(blog.TypeId)); blogType != nil {
		self.Data["blogType"] = blogType
	}

	setPreNextBlog(self, blog.Id, true)
	setPreNextBlog(self, blog.Id, false)

	go upateLookCount(blog.Id)

	self.displayWeb()
}

func upateLookCount(id int) {
	blog := &Blog{}
	blog.UpdateLookCount(id)
}

func setPreNextBlog(self *WebController, id int, pre bool) {
	key := "pre"
	condition := "id__gt"
	orderBy := "id"
	blog := &Blog{}
	fileters := make([]interface{}, 0)
	fileters = append(fileters, "status", 1)
	if !pre {
		key = "next"
		condition = "id__lt"
		orderBy = "-id"
	}
	fileters = append(fileters, condition, id)
	blogs := make([]*Blog, 0)
	blog.GetTopList(blog.TableName(), 1, &blogs, orderBy, fileters...)
	if len(blogs) > 0 {
		self.Data[key] = true
		self.Data[key+"Blog"] = blogs[0]
	} else {
		self.Data[key] = false
	}
}

func setRecentBlog(self *WebController) {
	blog := &Blog{}
	blogFileters := make([]interface{}, 0)
	blogFileters = append(blogFileters, "status", 1)
	blogs := make([]*Blog, 0)
	blog.GetTopList(blog.TableName(), 10, &blogs, "-id", blogFileters...)
	self.Data["recentBlogs"] = blogs
}

func setHotBlog(self *WebController) {
	blog := &Blog{}
	blogFileters := make([]interface{}, 0)
	blogFileters = append(blogFileters, "status", 1)
	blogs := make([]*Blog, 0)
	blog.GetTopList(blog.TableName(), 10, &blogs, "-look_count", blogFileters...)
	self.Data["hotBlogs"] = blogs
}

func setHomeData(self *WebController) {
	self.Data["blogTypes"] = getAllTypes()
	self.Data["blogTags"] = getAllTags()
	self.Data["config"] = getAllSysConfig()
	self.Data["blogLinks"] = getAllHomeBlogLink()

	setRecentBlog(self)
	setHotBlog(self)

}

func (self *WebController) Links() {
	self.Data["config"] = getAllSysConfig()
	self.Data["blogTypes"] = getAllTypes()
	self.Data["blogLinks"] = getAllBlogLink()
	self.displayWeb()
}
