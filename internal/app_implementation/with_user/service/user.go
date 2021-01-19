package service

import (
	"encoding/json"
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
	//3.处理下行
	t, _ := r.Cookie("token")
	u, _ := r.Cookie("uid")
	loginOutDeleteCookie(w, t)
	loginOutDeleteCookie(w, u)
	w.WriteHeader(http.StatusResetContent)
}

//NewPassword 修改密码
func NewPassword(w http.ResponseWriter, r *http.Request) {

	//获得参数
	req := struct_logic.ChangePassword{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	u := realize_logic.User{}
	uid, _ := r.Cookie("uid")
	u.UserID = uid.Value
	//处理参数
	if err := u.MyPassword(req.Password); err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic() //处理失败
	}

	//下行结果
	w.WriteHeader(http.StatusNoContent)
}

//MyIdentity 账号自身信息
func MyIdentity(w http.ResponseWriter, r *http.Request) {
	//获取账号ID
	u, _ := r.Cookie("uid")
	user := realize_logic.User{}
	user.UserID = u.Value
	user.CurrentUserInformation()

}
