# 抖声
呆头鹅大队项目——抖声（客户端已提供），主要使用go语言来完善相应的接口，尽可能提高性能。
# 运行方式
1. go mod init
2. go mod tidy
3. 修改配置文件，连接数据库
4. go run main.go router.go

# 数据库E-R图
![dousheng](https://wuhlan3-1307602190.cos.ap-guangzhou.myqcloud.com/img/dousheng.svg)

# 项目亮点
1. 采用repository、service、controller三层结构，结构清晰，模块之间耦合性较低
2. 使用JWT鉴权，直接对token进行解析，可以得到user_id，减少访问数据库的次数
3. 对密码进行加密，确保数据的安全性
# 优化方向
1. 加入redis来减少访问数据库的次数
2. 使用消息队列，来实现系统解耦和流量削峰
3. 将视频、封面等资源存放到COS存储桶中，便于管理，提高传输效率
4. 使用ffmpeg来获取视频封面
