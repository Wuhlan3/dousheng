# 抖声
呆头鹅大队项目——抖声（客户端已提供），主要使用go语言来完善相应的接口，尽可能提高性能。
# 运行方式
1. go mod init
2. go mod tidy
3. 修改配置文件，连接数据库
4. go run main.go router.go

``` bash
.
├── branchtest.txt
├── config             #文件配置
│   ├── config.go
│   └── config.yml
├── controller         #controller层
│   ├── comment.go
│   ├── common.go
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── doushengdb.sql      #创建数据库
├── go.mod
├── go.sum
├── LICENSE
├── log                 #日志文件
│   └── log.txt
├── main.go             #main
├── middleware          #中间件
│   └── jwt.go
├── public              #视频等资源
│   ├── bear.mp4
│   └── data
├── README.md
├── repository          #repository层
│   ├── db_init.go
│   └── user.go
├── router.go
├── service             #service层
│   ├── user_info.go        #用户信息
│   ├── user_login.go       #用户登录
│   └── user_register.go    #用户注册
└── util                    #通用工具
    └── logger.go           #日志系统配置
```