package user_web

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Logruser 中间件
func Logruser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// 执行前
		next.ServeHTTP(w, r) // web程序执行
		// 执行后
		endTime := time.Since(start)
		RequestTheAddress := r.RequestURI
		RequestTheHost := r.RemoteAddr
		logrus.Printf("Request The Address =%s，Request The Host =%s，This http  request dispose use time is %s", RequestTheAddress, RequestTheHost, endTime)
	})

}
