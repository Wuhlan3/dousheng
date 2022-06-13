# 抖声
呆头鹅大队项目——抖声（客户端已提供），主要使用go语言来完善相应的接口，尽可能提高性能。
# 运行方式
1. go mod init
2. go mod tidy
3. 修改配置文件，连接数据库
4. go run main.go router.go

# 目录结构

``` bash
.
├── branchtest.txt
├── config                  #配置文件
│   ├── config.go
│   └── config.yml
├── controller              #controller层
│   ├── comment.go
│   ├── common.go
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── doushengdb.sql          #DDL
├── go.mod
├── go.sum
├── LICENSE
├── log
│   └── log.txt
├── main.go     
├── middleware              #中间件
│   ├── AuthMiddleware.go
│   └── jwt.go
├── proto                   #proto相关结构体
│   ├── douyin_comment_list.proto
│   ├── douyin_publish_action.proto
│   ├── douyin_publish_list.proto
│   ├── douyin_relation_action.proto
│   ├── douyin_relation_follower_list.proto
│   ├── douyin_relation_follow_list.proto
│   ├── feed.proto
│   └── proto
│       ├── douyin_comment_list.pb.go
│       ├── douyin_publish_action.pb.go
│       ├── douyin_publish_list.pb.go
│       ├── douyin_relation_action.pb.go
│       ├── douyin_relation_follower_list.pb.go
│       ├── douyin_relation_follow_list.pb.go
│       └── feed.pb.go
├── public                  #静态资源
│   ├── 4_jiajia.jpg
│   ├── 4_jiajia.mp4
│   ├── bear.jpg
│   ├── bear.mp4
│   ├── bear.png
│   ├── cat1.jpg
│   ├── cat1.mp4
│   ├── cat2.jpg
│   ├── cat2.mp4
│   └── data
├── README.md
├── repository              #repository层
│   ├── comment.go
│   ├── db_init.go
│   ├── favourite.go
│   ├── relation.go
│   ├── user.go
│   ├── video.go
│   └── videos.go
├── router.go               #路由
├── service                 #service层
│   ├── comment_action.go
│   ├── comment_list.go
│   ├── favourite_action.go
│   ├── favourite_list.go
│   ├── feed.go
│   ├── publish.go
│   ├── publish_list.go
│   ├── relation_action.go
│   ├── relation_follower_list.go
│   ├── relation_follow_list.go
│   ├── user_info.go
│   ├── user_login.go
│   └── user_register.go
└── util                    #日志系统
    └── logger.go
```

# 项目亮点
1. 采用repository、service、controller三层结构，结构清晰，模块之间耦合性较低
2. 使用JWT鉴权，直接对token进行解析，可以得到user_id，减少访问数据库的次数
3. 对密码进行加密，确保数据的安全性
# 优化方向
1. 加入redis来减少访问数据库的次数
2. 使用消息队列，来实现系统解耦和流量削峰
3. 将视频、封面等资源存放到COS存储桶中，便于管理，提高传输效率
4. 使用ffmpeg来获取视频封面
