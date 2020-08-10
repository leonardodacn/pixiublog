package main

import (
	"pixiublog/jobs"
	"pixiublog/models"
	_ "pixiublog/routers"
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
	beego.Run()
}
