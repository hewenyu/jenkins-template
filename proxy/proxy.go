package proxy

import "github.com/gin-gonic/gin"


/*
InitForward 初始化路由转发模块
*/
func InitForward(Router *gin.RouterGroup) {
	
	service := NewDefaultProxy("https://www.baidu.com", "/")
	
	service.Gin(Router)
}