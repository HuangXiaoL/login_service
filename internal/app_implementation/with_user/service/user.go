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
	loginSuccessfullyIssuedCookie(w, "token", token, 0)
	loginSuccessfullyIssuedCookie(w, "uid", u.UserID, 0)
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
	u := &realize_logic.User{}
	token, err := u.Login(req, 0)
	if err != nil { // 登录失败。
		logrus.Println(err)
		httpkit.WrapError(err).WithStatus(http.StatusUnauthorized).Panic()
	}
	loginSuccessfullyIssuedCookie(w, "token", token, req.Remember)
	loginSuccessfullyIssuedCookie(w, "uid", u.UserID, req.Remember)
	w.WriteHeader(http.StatusCreated)
}

//UserLoginOut 用户退出登录
func UserLoginOut(w http.ResponseWriter, r *http.Request) {
	t, _ := r.Cookie("token")
	u, _ := r.Cookie("uid")
	us := &realize_logic.User{}
	logrus.Println(t.Value)
	logrus.Println(u.Value)
	us.UserID = u.Value
	if err := us.LoginOut(t.Value); err != nil {
		logrus.Info(err)
		httpkit.WrapError(err).WithStatus(http.StatusForbidden).Panic()
	}
	logrus.Info(321312312312)
	loginOutDeleteCookie(w, t)
	loginOutDeleteCookie(w, u)
	w.WriteHeader(http.StatusResetContent)
}

//loginSuccessfullyIssuedCookie 登录成功下发cookie数据
func loginSuccessfullyIssuedCookie(w http.ResponseWriter, name string, value string, remember int) {
	maxAge := int(60 * time.Second)
	if remember == 1 {
		maxAge = COOKIE_MAX_MAX_AGE // 七天
	}
	//3.结果下行
	cookie := &http.Cookie{
		Name:   name,
		Value:  value,
		MaxAge: maxAge,
	}
	http.SetCookie(w, cookie)
}

//loginOutDeleteCookie 退出登录删除cookie
func loginOutDeleteCookie(w http.ResponseWriter, cookie *http.Cookie) {
	cookie.Value = ""
	cookie.MaxAge = 0
	http.SetCookie(w, cookie)
}
