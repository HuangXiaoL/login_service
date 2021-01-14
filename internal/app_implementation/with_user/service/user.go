package service

import (
	"net/http"
	"time"

	realize_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/logic"

	"github.com/sirupsen/logrus"

	struct_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/logic"

	"github.com/joyparty/httpkit"
)

const COOKIE_MAX_MAX_AGE = int(time.Hour * 24 * 7 / time.Second) //7*24 小时 单位：秒。
//RegisterUserInfo 注册用户
func RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	//1.上传参数赋值
	req := struct_logic.Register{}
	_ = r.ParseForm()
	if err := httpkit.ScanValues(&req, r.PostForm); err != nil { //绑定参数
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	//2.数据处理
	//执行注册流程
	u := realize_logic.User{}
	if err := u.RegisterInfo(req); err != nil { //注册失败，uuid或者email重复 ，证明该邮箱已经被注册了，下行错误代码 409
		logrus.Println(err)
		httpkit.WrapError(err).WithStatus(http.StatusConflict).Panic()
	}
	//执行登录流程,生成session salt，下行201
	l := struct_logic.Login{}
	l.Email = u.Email
	token, err := u.Login(l, 1) //way 参数为1 不需要效验账号密码，直接登录
	if err != nil {             //注册成功 登录失败。返回200 跳转到登录接口进行登录
		logrus.Println(err)
		httpkit.WrapError(err).WithStatus(http.StatusOK).Panic()
	}
	//3.结果下行
	// 注册成功，登录成功，下发cookie
	maxAge := int(60 * 15 * time.Second) //15分钟
	uc := &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: maxAge,
	}
	http.SetCookie(w, uc)
	w.WriteHeader(http.StatusCreated)
}

//UserLogin 用户登录
func UserLogin(w http.ResponseWriter, r *http.Request) {
	//1.上传参数赋值
	req := struct_logic.Login{}
	_ = r.ParseForm()
	if err := httpkit.ScanValues(&req, r.PostForm); err != nil { //绑定参数
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	//2.数据处理
	logrus.Printf("%+v", req)
	u := realize_logic.User{}
	token, err := u.Login(req, 0)
	if err != nil { // 登录失败。
		logrus.Println(err)
		httpkit.WrapError(err).WithStatus(http.StatusUnauthorized).Panic()
	}
	maxAge := int(60 * time.Second)
	if req.Remember == 1 {
		maxAge = COOKIE_MAX_MAX_AGE // 七天
	}
	//3.结果下行
	uc := &http.Cookie{
		Name:   "token",
		Value:  token,
		MaxAge: maxAge,
	}
	http.SetCookie(w, uc)
	w.WriteHeader(http.StatusCreated)
}
