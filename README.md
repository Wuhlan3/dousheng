# 抖声
呆头鹅大队项目——抖声（客户端已提供），主要使用go语言来完善相应的接口，尽可能提高性能。
也可以参考一下，微服务结构版本https://github.com/Wuhlan3/kitexdousheng

# 运行方式
1. go mod init
2. go mod tidy
3. 修改配置文件，连接数据库
4. go run main.go router.go

# 数据库E-R图
![dousheng](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng.svg)

# feed过程
feed即用户在刷视频过程中请求的接口，响应的是视频相关数据，这一部分应该是最频繁调用的且包括了几乎所有表的数据，所以该过程较复杂。
1. 用户会请求两个参数，分别是token和latest_time。其中token会经过JWT解析，得到用户的uid，latest_time表示限制返回视频的时间戳；
2. 由于需要限制返回的视频数量，且我们期望能够优先刷到最新投稿的视频，所以可以采用Redis中的ZSET数据结构来保存视频的序列号；
3. 为了减少视频信息的查询数据库次数，当我们获得视频序列号的时候，可以直接通过video:id在Redis中查询相应的视频信息。
其流程图如下：

<img src="https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng_feed.jpg" width="700px">

# 项目亮点
1. 采用repository、service、controller三层结构，结构清晰，模块之间耦合性较低
2. 使用JWT鉴权，直接对token进行解析，可以得到user_id，减少访问数据库的次数
3. 对密码进行加密，确保数据的安全性
4. 将视频、封面等资源存放到COS存储桶中，便于管理，提高传输效率
5. 加入redis来减少访问数据库的次数
6. 使用ffmpeg来获取视频封面

# 优化方向
1. 使用消息队列，来实现系统解耦和流量削峰
2. 完善数据校验过程

# 运行结果
注册与登录、视频流功能如下：

![dousheng_result1](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng1.png)

点赞、关注、喜欢视频列表、评论等功能如下：

![dousheng_result2](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng2.png)
