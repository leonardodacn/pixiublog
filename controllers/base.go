package controllers

import (
	"pixiublog/models"
	"pixiublog/utils"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

const (
	MSG_OK  = 0
	MSG_ERR = -1
)

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	user           *models.Admin
	userId         int
	userName       string
	loginName      string
	pageSize       int
	allowUrl       string
	serverGroups   string
	taskGroups     string
}

// GetString returns the input value by key string or the default value while it's present and input is blank
func (c *BaseController) GetString(key string, def ...string) string {
	if v := c.Ctx.Input.Query(key); v != "" {
		return strings.TrimSpace(v)
	}
	if len(def) > 0 {
		return def[0]
	}
	return ""
}

// GetAlwaysInt returns input as an int or the default value while it's present and input is blank, if errors , return 0
func (c *BaseController) GetAlwaysInt(key string, def ...int) int {
	strv := c.Ctx.Input.Query(key)
	if len(strv) == 0 && len(def) > 0 {
		return def[0]
	}

	if v, err := strconv.Atoi(strv); err == nil {
		return v
	}

	if len(def) > 0 {
		return def[0]
	}

	return 0
}

//前期准备
func (self *BaseController) Prepare() {
	self.pageSize = 20
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["version"] = beego.AppConfig.String("version")
	self.Data["siteName"] = beego.AppConfig.String("site.name")
	self.Data["curRoute"] = self.controllerName + "." + self.actionName
	self.Data["curController"] = self.controllerName
	self.Data["curAction"] = self.actionName

	self.Auth()
	self.Data["loginUserId"] = self.userId
	self.Data["loginUserName"] = self.userName
}

//登录权限验证
func (self *BaseController) Auth() {
	arr := strings.Split(self.Ctx.GetCookie("auth"), "|")
	self.userId = 0
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userId, _ := strconv.Atoi(idstr)
		if userId > 0 {
			user, err := models.AdminGetById(userId)

			//if err == nil && password == utils.Md5([]byte(self.getClientIp()+"|"+user.Password+user.Salt)) {
			if err == nil && password == utils.Md5([]byte(user.Password+user.Salt)) {
				self.userId = user.Id
				self.loginName = user.LoginName
				self.userName = user.RealName
				self.user = user
				self.AdminAuth()
				self.dataAuth(user)
			}

			isHasAuth := strings.Contains(self.allowUrl, self.controllerName+"/"+self.actionName)
			excludeAction := "/loginin/loginout"
			excludeController := "/web"
			isNoAuth := strings.Contains(excludeAction, self.actionName) || strings.Contains(excludeController, self.controllerName)

			if isHasAuth == false && isNoAuth == false {
				if strings.Contains(self.actionName, "ajax") {
					self.ajaxMsg("没有权限", MSG_ERR)
					return
				}

				flash := beego.NewFlash()
				flash.Error("没有权限")
				flash.Store(&self.Controller)
				return
			}
		}
	}

	if self.userId == 0 &&
		(self.controllerName != "webseo" && self.controllerName != "web" && self.controllerName != "login" && self.actionName != "loginin") {
		self.redirect(beego.URLFor("LoginController.Login"))
	}
}

func (self *BaseController) dataAuth(user *models.Admin) {
	if user.RoleIds == "0" || user.Id == 1 {
		return
	}

	Filters := make([]interface{}, 0)
	Filters = append(Filters, "status", 1)

	RoleIdsArr := strings.Split(user.RoleIds, ",")

	RoleIds := make([]int, 0)
	for _, v := range RoleIdsArr {
		id, _ := strconv.Atoi(v)
		RoleIds = append(RoleIds, id)
	}

	Filters = append(Filters, "id__in", RoleIds)

	Result, _ := models.RoleGetList(1, 1000, Filters...)
	serverGroups := ""
	taskGroups := ""
	for _, v := range Result {
		serverGroups += v.ServerGroupIds + ","
		taskGroups += v.TaskGroupIds + ","
	}

	self.serverGroups = strings.Trim(serverGroups, ",")
	self.taskGroups = strings.Trim(taskGroups, ",")
}

func (self *BaseController) AdminAuth() {
	// 左侧导航栏
	filters := make([]interface{}, 0)
	filters = append(filters, "status", 1)
	if self.userId != 1 {
		//普通管理员
		adminAuthIds, _ := models.RoleAuthGetByIds(self.user.RoleIds)
		adminAuthIdArr := strings.Split(adminAuthIds, ",")
		filters = append(filters, "id__in", adminAuthIdArr)
	}
	result, _ := models.AuthGetList(1, 1000, filters...)
	list := make([]map[string]interface{}, len(result))
	list2 := make([]map[string]interface{}, len(result))
	allow_url := ""
	i, j := 0, 0
	for _, v := range result {
		if v.AuthUrl != " " || v.AuthUrl != "/" {
			allow_url += v.AuthUrl
		}
		row := make(map[string]interface{})
		if v.Pid == 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list[i] = row
			i++
		}
		if v.Pid != 1 && v.IsShow == 1 {
			row["Id"] = int(v.Id)
			row["Sort"] = v.Sort
			row["AuthName"] = v.AuthName
			row["AuthUrl"] = v.AuthUrl
			row["Icon"] = v.Icon
			row["Pid"] = int(v.Pid)
			list2[j] = row
			j++
		}
	}

	self.Data["SideMenu1"] = list[:i]  //一级菜单
	self.Data["SideMenu2"] = list2[:j] //二级菜单

	self.allowUrl = allow_url + "/home/index"
}

// 是否POST提交
func (self *BaseController) isPost() bool {
	return self.Ctx.Request.Method == "POST"
}

//获取用户IP地址
func (self *BaseController) getClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}

// 重定向
func (self *BaseController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}

//加载模板
func (self *BaseController) display(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

//加载模板
func (self *BaseController) displaySeo(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "tpl"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".tpl"
	}
	self.TplName = tplname
}

var currentYear = beego.Date(time.Now(), "Y")

//加载模板
func (self *BaseController) displayWeb(tpl ...string) {
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "web/web.html"
	self.Data["currentYear"] = currentYear
	self.TplName = tplname
}

//ajax返回
func (self *BaseController) ajaxMsg(msg interface{}, msgno int) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["message"] = msg
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//ajax返回 列表
func (self *BaseController) ajaxList(msg interface{}, msgno int, count int64, data interface{}) {
	out := make(map[string]interface{})
	out["code"] = msgno
	out["msg"] = msg
	out["count"] = count
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
	self.StopRun()
}

//资源分组信息
func serverGroupLists(authStr string, adminId int) (sgl map[int]string) {
	Filters := make([]interface{}, 0)
	Filters = append(Filters, "status", 1)
	if authStr != "0" && adminId != 1 {
		serverGroupIdsArr := strings.Split(authStr, ",")
		serverGroupIds := make([]int, 0)
		for _, v := range serverGroupIdsArr {
			id, _ := strconv.Atoi(v)
			serverGroupIds = append(serverGroupIds, id)
		}
		Filters = append(Filters, "id__in", serverGroupIds)
	}

	groupResult, n := models.ServerGroupGetList(1, 1000, Filters...)
	sgl = make(map[int]string, n)
	for _, gv := range groupResult {
		sgl[gv.Id] = gv.GroupName
	}
	//sgl[0] = "默认分组"
	return sgl
}

func taskGroupLists(authStr string, adminId int) (gl map[int]string) {
	groupFilters := make([]interface{}, 0)
	groupFilters = append(groupFilters, "status", 1)
	if authStr != "0" && adminId != 1 {
		taskGroupIdsArr := strings.Split(authStr, ",")
		taskGroupIds := make([]int, 0)
		for _, v := range taskGroupIdsArr {
			id, _ := strconv.Atoi(v)
			taskGroupIds = append(taskGroupIds, id)
		}
		groupFilters = append(groupFilters, "id__in", taskGroupIds)
	}
	groupResult, n := models.GroupGetList(1, 1000, groupFilters...)
	gl = make(map[int]string, n)
	for _, gv := range groupResult {
		gl[gv.Id] = gv.GroupName
	}
	return gl
}

func serverListByGroupId(groupId int) []string {
	Filters := make([]interface{}, 0)
	Filters = append(Filters, "status", 1)
	Filters = append(Filters, "group_id", groupId)
	Result, _ := models.TaskServerGetList(1, 1000, Filters...)

	servers := make([]string, 0)
	for _, v := range Result {
		servers = append(servers, strconv.Itoa(v.Id), v.ServerName)
	}

	return servers
}

type AdminInfo struct {
	Id       int
	Email    string
	Phone    string
	RealName string
}

func AllAdminInfo(adminIds string) []*AdminInfo {
	Filters := make([]interface{}, 0)
	Filters = append(Filters, "status", 1)
	//Filters = append(Filters, "id__gt", 1)
	var notifyUserIds []int
	if adminIds != "0" && adminIds != "" {
		notifyUserIdsStr := strings.Split(adminIds, ",")
		for _, v := range notifyUserIdsStr {
			i, _ := strconv.Atoi(v)
			notifyUserIds = append(notifyUserIds, i)
		}
		Filters = append(Filters, "id__in", notifyUserIds)
	}
	Result, _ := models.AdminGetList(1, 1000, Filters...)

	adminInfos := make([]*AdminInfo, 0)
	for _, v := range Result {
		ai := AdminInfo{
			Id:       v.Id,
			Email:    v.Email,
			Phone:    v.Phone,
			RealName: v.RealName,
		}
		adminInfos = append(adminInfos, &ai)
	}

	return adminInfos
}

type serverList struct {
	GroupId   int
	GroupName string
	Servers   map[int]string
}

func serverLists(authStr string, adminId int) (sls []serverList) {
	serverGroup := serverGroupLists(authStr, adminId)
	Filters := make([]interface{}, 0)
	Filters = append(Filters, "status__in", []int{0, 1})

	Result, _ := models.TaskServerGetList(1, 1000, Filters...)
	for k, v := range serverGroup {
		sl := serverList{}
		sl.GroupId = k
		sl.GroupName = v
		servers := make(map[int]string)
		for _, sv := range Result {
			if sv.GroupId == k {
				servers[sv.Id] = sv.ServerName
			}
		}
		sl.Servers = servers
		sls = append(sls, sl)
	}
	return sls
}
