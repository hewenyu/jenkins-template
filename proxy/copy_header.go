package proxy

import (
	"github.com/hewenyu/jenkins-template/logger"
	"net/http"
	"os"
)

func CopyHeader(source http.Header, dest *http.Header) {
	for n, v := range source {
		for _, vv := range v {
			dest.Add(n, vv)
		}
	}
}

func fatal(err error) {
	if err != nil {
		logger.CLog.Info(err.Error())
		os.Exit(1)
	}
}
