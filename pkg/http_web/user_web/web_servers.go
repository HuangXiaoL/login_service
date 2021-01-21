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

	r.With(Logruser).Post(`/register`, service.RegisterUserInfo)      //注册账号
	r.With(Logruser).Post(`/login`, service.UserLogin)                // 账号登录
	r.With(LoginAuth).Post(`/my/password`, service.NewPassword)       //修改密码
	r.With(LoginAuth).Post(`/user/:userID/lock`, service.NewPassword) //修改密码

	r.With(LoginAuth).Delete(`/login`, service.UserLoginOut) // 退出登录

	r.Route("/user/{userID}", func(r chi.Router) {
		r.Use(AdminAccessLevel)
		r.Post("/lock", service.LockUser)              // POST /user/123/lock  锁定账号
		r.Delete("/lock", service.UnLockUser)          // Delete /user/123/lock  解除锁定
		r.Put("/role", service.SetTheRole)             // Put /user/123/role  定义角色
		r.Delete("/password", service.DefaultPassword) // Delete user/123/password  重置密码
	})

	return r
}
