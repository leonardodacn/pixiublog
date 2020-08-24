package models

import (
	"net/url"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var StartTime int64

func Init(startTime int64) {
	StartTime = startTime
	dbhost := beego.AppConfig.String("db.host")
	dbport := beego.AppConfig.String("db.port")
	dbuser := beego.AppConfig.String("db.user")
	dbpassword := beego.AppConfig.String("db.password")
	dbname := beego.AppConfig.String("db.name")
	timezone := beego.AppConfig.String("db.timezone")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(
		new(Admin),
		new(Auth),
		new(Role),
		new(RoleAuth),
		new(ServerGroup),
		new(TaskServer),
		new(Ban),
		new(Group),
		new(Task),
		new(TaskLog),
		new(NotifyTpl),
		new(SysConfig),
		new(BlogType),
		new(BlogTag),
		new(BlogLink),
		new(BlogComment),
		new(Blog),
		new(BlogTagRelation),
	)

	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
	}
}

func TableName(name string) string {
	return beego.AppConfig.String("db.prefix") + name
}
