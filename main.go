package main

import (
	"pixiublog/jobs"
	"pixiublog/models"
	_ "pixiublog/routers"
	"pixiublog/utils"
	"time"

	"github.com/astaxie/beego"
)

func init() {

	//初始化数据模型
	var StartTime = time.Now().Unix()
	models.Init(StartTime)
	jobs.InitJobs()
}

func main() {
	//添加自定义的模板函数
	addTemplateFunc()
	beego.Run()
}

func addTemplateFunc() {
	beego.AddFuncMap("NumEq", utils.NumEq)
	beego.AddFuncMap("ModEq", utils.ModEq)
	beego.AddFuncMap("ContainNum", utils.ContainNum)
	beego.AddFuncMap("RandNum", utils.RandNum)
	beego.AddFuncMap("Add", utils.Add)
}
