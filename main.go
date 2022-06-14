package main

import (
	"dousheng/config"
	"dousheng/repository"
	"dousheng/util"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	config.InitConfig() // 设置配置文件
	if err := repository.Init(); err != nil {
		os.Exit(-1)
	} //数据库连接
	util.InitLogger()
	r := gin.Default()

	initRouter(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
