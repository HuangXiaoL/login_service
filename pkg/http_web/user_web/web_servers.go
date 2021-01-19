package user_web

import (
	"github.com/joyparty/httpkit"
	"github.com/sirupsen/logrus"
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/service"

	"github.com/go-chi/chi"
)

func NewRouter() *chi.Mux {
	r := chi.NewMux()

	hl := logrus.WithField("@type", "http")
	r.Use(httpkit.LogRequest(hl))
	r.Use(httpkit.Recoverer(hl))
	//接口
	r.With(LoginAuth).Get(`/my/identity`, service.MyIdentity) // 本账号信息接口

	r.With(Logruser).Post(`/register`, service.RegisterUserInfo) //注册账号
	r.With(Logruser).Post(`/login`, service.UserLogin)           // 账号登录
	r.With(LoginAuth).Post(`/my/password`, service.NewPassword)  //修改密码

	r.With(LoginAuth).Delete(`/login`, service.UserLoginOut) // 退出登录

	return r
}
