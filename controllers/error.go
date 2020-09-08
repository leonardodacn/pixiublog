package controllers

import "github.com/astaxie/beego"

//该控制器处理页面错误请求
type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error401() {
	c.TplName = "error/401.html"
	displayError(c)
}

func (c *ErrorController) Error403() {
	c.TplName = "error/403.html"
	displayError(c)
}

func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
	displayError(c)
}

func (c *ErrorController) Error500() {
	c.Data["content"] = "server error"
	c.TplName = "error/500.html"
	displayError(c)
}
func (c *ErrorController) Error503() {
	c.TplName = "error/503.html"
	displayError(c)
}

func displayError(c *ErrorController) {
	c.Layout = "web/web.html"
	c.Data["blogTypes"] = getAllTypes()
}
