# pixiublog

## 一、说明

pixiublog是一款轻量级博客系统，go语言开发，部署超级简单，资源消耗少，运行稳定。

### 使用的框架

+ beego
+ bee
+ layui
+ bootstrap
+ PPGO_JOB

## 二、安装方法

### 方法一、 编译安装

- go get https://github.com/leonardodacn/pixiublog
- 创建mysql数据库，并将pixiublog.sql导入
- 修改config 配置数据库
- 运行 go build
- 运行 ./run.sh start|stop

mac
- 运行 ./package.sh -a amd64 -p darwin -v v2.x.0

linux
- 运行 ./package.sh -a 386 -p linux -v v2.x.0
- 运行 ./package.sh -a amd64 -p linux -v v2.x.0

windows
- 运行 ./package.sh -a amd64 -p windows -v v2.x.0


### 方法二、直接使用

#### linux

- 进入 https://pixiublog/releases
- 下载 pixiublog-linux-2.x.0.zip 并解压
- 进入文件夹，设置好数据库(创建数据库，导入pixiublog.sql)和配置文件(conf/app.conf)
- 运行 ./run.sh start|stop

#### mac

- 进入https://pixiublog/releases
- 下载 pixiublog-mac-2.x.0.zip 并解压
- 进入文件夹，设置好数据库(创建数据库，导入pixiublog.sql)和配置文件(conf/app.conf)
- 运行 ./run.sh start|stop

#### windows

- 进入 https://pixiublog/releases
- 下载 pixiublog-windows-2.x.0.zip 并解压
- 进入文件夹，设置好数据库(创建数据库，导入pixiublog.sql)和配置文件(conf/app.conf)
- 运行 run.bat

### 访问
+ 前台访问：http://your_host:port
+ 后台访问：http://your_host:port/admin
+ 用户名：admin 密码：123456

### 配置文件
根据自己的情况修改数据库和启动端口
```
appname = pixiublog
httpport = 8080
runmode = dev

version= V1.0

# 允许同时运行的任务数
jobs.pool = 1000

# 站点名称
site.name = 博客后台管理器

#通知方式 0=邮件，1=信息，2=钉钉，3=微信
notify.type = 0

# 数据库配置
db.host = 127.0.0.1
db.user = root
db.password = "123456"
db.port = 3306
db.name = pixiublog
db.prefix = 
db.timezone = Asia/Shanghai

# 邮件通知配置
email.host = smtp.xxx.com
email.port = 25
email.from = ci@xxx.cn
email.user = ci@xxx.cn
email.password = "xxxxxx"
email.pool = 10

# 短信通知方式配置
msg.url = http://xxx.com/api/tools/send_sms
msg.pool = 10

# 钉钉通知配置
dingtalk.url = "https://oapi.dingtalk.com/robot/send?access_token=%s"
dingtalk.pool = 10

# 微信通知方式配置
wechat.url = http://xx.com/api/tools/send_wechat
wechat.pool = 10
```

## 三、编译安装-可能会遇到的问题
----
go build 时遇到以下错误：
jobs/job.go:19:2: cannot find package "golang.org/x/crypto/ssh" in any of:

需要 git clone https://github.com/golang/crypto.git
并拷贝到 $GOPATH/src/golang.org/x/ 下就OK

或

git clone https://github.com/golang/crypto.git $GOPATH/src/golang.org/x/crypto

## 四、Docker

本地编译好的二进制文件放在根目录下执行下面的命令即可拥有
```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
```

```
docker-compose up -d

#日志查看
docker-compose logs -f web

```

## 五、其它

### 生成代码命令
`bee generate appcode -tables="sys_config" -driver=mysql -conn="root:123456@tcp(127.0.0.1:3306)/pixiublog" -level=3
`