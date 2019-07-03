package app

import (
	"github.com/chulinshao/rehab/web/middleware"
	"github.com/chulinshao/rehab/web/routers"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	time.LoadLocation("Asia/Shanghai")
}

func StartApp() {
	router := gin.Default()
	routers.ConfigRouter(router)
	router.Use(middleware.GlobalOriginMiddleware())
	router.Run(":8080")
}
