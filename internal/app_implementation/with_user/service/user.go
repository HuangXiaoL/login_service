package service

import (
	"encoding/json"
	"net/http"
	"time"

	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/custom_error"

	"github.com/go-chi/chi"

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
		if err.Error() == custom_error.USER_LOCK {
			logrus.Info(err)
			httpkit.WrapError(err).WithStatus(http.StatusLocked).Panic()
		}
		logrus.Info(err)
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
	result, err := user.CurrentUserInformation()
	if err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic() //处理失败
	}
	buf, err := json.MarshalIndent(result, "", "    ") //格式化编码
	if err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic() //处理失败
	}
	httpkit.Render.String(w, 200, string(buf))
}

//LockUser 锁定账号
func LockUser(w http.ResponseWriter, r *http.Request) {
	//获取数据，参数效验
	lockID := chi.URLParam(r, "userID")
	//锁定处理
	us := realize_logic.User{}
	if err := us.LockTheAccount(lockID); err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusInternalServerError).Panic()
	}
	//下行结果
	w.WriteHeader(http.StatusResetContent)
}

//UnLockUser 解锁账号
func UnLockUser(w http.ResponseWriter, r *http.Request) {
	lockID := chi.URLParam(r, "userID")
	u, _ := r.Cookie("uid")
	user := realize_logic.User{}
	user.UserID = u.Value
	result, err := user.CurrentUserInformation()
	if err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusForbidden).Panic()
	}
	if result.Role != "admin" {
		httpkit.WrapError(err).WithStatus(http.StatusForbidden).Panic()
	}
	if lockID == result.ID {
		httpkit.WrapError(err).WithStatus(http.StatusNotAcceptable).Panic() //不允许自己解锁自己
	}
	//锁定处理
	us := realize_logic.User{}
	if err := us.UNLockTheAccount(lockID); err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusInternalServerError).Panic()
	}
	//下行结果
	w.WriteHeader(http.StatusResetContent)
}

//SetTheRole 设置角色
func SetTheRole(w http.ResponseWriter, r *http.Request) {
	// 参数获取
	lockID := chi.URLParam(r, "userID") //被设置的用户ID
	req := struct_logic.SetRole{}       //设置的角色参数
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	//数据处理
	user := realize_logic.User{}
	if err := user.SetUserRole(lockID, req.Role); err != nil {
		if err.Error() == custom_error.NO_USER { //查询用户失败了
			logrus.Println(err)
			httpkit.WrapError(err).WithStatus(http.StatusNotFound).Panic()
		}
		if err.Error() == custom_error.NO_ROLE { //查询角色失败了
			logrus.Println(err)
			httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic()
		}
		httpkit.WrapError(err).WithStatus(http.StatusNotFound).Panic()
	}
	//下行结果
	w.WriteHeader(http.StatusNoContent)
}

//DefaultPassword 设置为默认的密码---重置密码
func DefaultPassword(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	lockID := chi.URLParam(r, "userID") //被设置的用户ID
	//处理参数
	u := realize_logic.User{}
	u.UserID = lockID
	if err := u.MyPassword("123456"); err != nil {
		httpkit.WrapError(err).WithStatus(http.StatusBadRequest).Panic() //处理失败
	}
	//下行结果
	w.WriteHeader(http.StatusNoContent)
}
