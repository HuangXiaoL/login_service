package user_web

import (
	"net/http"
	"time"

	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/service"

	"github.com/joyparty/httpkit"

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
		logrus.Printf("Request The Address =%s，Request The Host =%s，This http  request dispose use time is %s，now time is %s", RequestTheAddress, RequestTheHost, endTime, time.Now())
	})

}

// Logruser 中间件 401 未登录，禁止匿名访问
func LoginAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t, errtoken := r.Cookie("token")
		u, erruid := r.Cookie("uid")
		if errtoken != nil || erruid != nil || len(t.Value) < 1 || len(u.Value) < 1 {
			httpkit.WrapError(errtoken).WithStatus(http.StatusUnauthorized).Panic()
		}
		if err := service.AuthenticationToken(r); err != nil {
			logrus.Info(err)
			httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic()
		}
		// 执行前
		next.ServeHTTP(w, r) // web程序执行
	})

}
