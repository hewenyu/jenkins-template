package jenkins_template

import (
	"fmt"
	"github.com/hewenyu/jenkins-template/logger"
	"github.com/hewenyu/jenkins-template/proxy"
	"syscall"
	"time"
	
	"github.com/fvbock/endless"
)

func main() {
	Router := proxy.Routers()
	
	endless.DefaultReadTimeOut = 10 * time.Millisecond
	endless.DefaultWriteTimeOut = 10 * time.Second
	endless.DefaultMaxHeaderBytes = 1 << 20
	
	endPoint := fmt.Sprintf(":%d", 8080)
	
	server := endless.NewServer(endPoint, Router)
	
	server.BeforeBegin = func(add string) {
		logger.CLog.Info(fmt.Sprintf("Actual pid is %d", syscall.Getpid()))
	}
	
	err := server.ListenAndServe()
	if err != nil {
		logger.CLog.Error(fmt.Sprintf("Server err: %v", err))
	}
}
