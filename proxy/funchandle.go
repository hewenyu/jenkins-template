package proxy

import (
	"github.com/gin-gonic/gin"
	"github.com/hewenyu/jenkins-template/logger"
	"io/ioutil"
	"net/http"
)

type DefaultProxy struct {
	Addr string // 地址
	Path string // 转发地址
}

func NewDefaultProxy(addr, path string) *DefaultProxy {
	
	return &DefaultProxy{
		Addr: addr,
		Path: path,
	}
}

/*
Gin 用于嵌入gin
*/
func (srv *DefaultProxy) Gin(router *gin.RouterGroup) {
	
	// AuthGroup := router.Group("")
	
	router.Any("/", gin.WrapF(srv.HandlerProxy))
	router.Any("/:path", gin.WrapF(srv.HandlerProxy))
}

/*
HandlerProxy 转发功能实现
*/
func (srv *DefaultProxy) HandlerProxy(w http.ResponseWriter, r *http.Request) {
	uri := srv.Addr + r.RequestURI
	
	logger.CLog.Info(r.Method + ": " + uri)
	
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)
		fatal(err)
		logger.CLog.Sugar().Infof("Body: %v\n", string(body))
	}
	
	rr, err := http.NewRequest(r.Method, uri, r.Body)
	
	fatal(err)
	
	CopyHeader(r.Header, &rr.Header)
	
	// 创建新的请求
	var transport http.Transport
	resp, err := transport.RoundTrip(rr)
	
	fatal(err)
	
	logger.CLog.Sugar().Infof("Resp-Headers: %v\n", resp.Header)
	
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	
	fatal(err)
	
	dH := w.Header()
	CopyHeader(resp.Header, &dH)
	dH.Add("Requested-Host", rr.Host)
	
	_, _ = w.Write(body)
}