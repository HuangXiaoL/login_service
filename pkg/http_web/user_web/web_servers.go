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
	r.With(Logruser).Post(`/register`, service.RegisterUserInfo)
	r.With(Logruser).Post(`/login`, service.UserLogin)
	r.With(LoginAuth).Delete(`/login`, service.UserLoginOut)

	return r
}
