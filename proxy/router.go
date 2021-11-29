package proxy

import "github.com/gin-gonic/gin"

// Routers 初始化总路由
func Routers() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	var DefultRouter = gin.Default()
	
	// gin.HandlerFunc()
	
	// 不需要验证的路由
	PublicGroup := DefultRouter.Group("")
	{
		InitForward(PublicGroup) // 转发路由模块
	}
	
	return DefultRouter
}
