package main

import (
	"github.com/chulinshao/rehab/web/router"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")
	router.ConfigUserRouter(app)
	//开启web服务
	app.Run(
		//开启web服务
		iris.Addr(":8080"),
		// 按下CTRL / CMD + C时跳过错误的服务器：
		iris.WithoutServerError(iris.ErrServerClosed),
		//实现更快的json序列化和更多优化：
		iris.WithOptimizations,
	)
}
