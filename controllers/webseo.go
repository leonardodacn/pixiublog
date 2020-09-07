package controllers

import (
	"time"

	"github.com/astaxie/beego"
)

type WebSeoController struct {
	BaseController
}

func (self *WebSeoController) Robots() {
	self.Data["config"] = getAllSysConfig()
	self.Ctx.Output.Header("Content-Type", "text/plain;charset=UTF-8")
	self.displaySeo()
}

func (self *WebSeoController) SiteMapHtml() {
	siteMap(self)
	self.displaySeo()
}

func (self *WebSeoController) SiteMapXml() {
	siteMap(self)
	self.Ctx.Output.Header("Content-Type", "application/xml;charset=UTF-8")
	self.displaySeo()
}

func (self *WebSeoController) SiteMapTxt() {
	siteMap(self)
	self.Ctx.Output.Header("Content-Type", "text/plain;charset=UTF-8")
	self.displaySeo()
}

func siteMap(self *WebSeoController) {
	var currentDate = beego.Date(time.Now(), "Y-m-d")
	self.Data["currentDate"] = currentDate
	self.Data["blogTypes"] = getAllTypes()
	self.Data["blogTags"] = getAllTags()
	self.Data["blogs"] = getAllBlogs()
	self.Data["config"] = getAllSysConfig()
}
