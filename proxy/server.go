package proxy

import (
	"github.com/hewenyu/jenkins-template/logger"
	"net/http"
	

)

func Server() {
	service := NewDefaultProxy("https://www.baidu.com", "/")
	http.HandleFunc("/", service.HandlerProxy)
	logger.CLog.Sugar().Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
