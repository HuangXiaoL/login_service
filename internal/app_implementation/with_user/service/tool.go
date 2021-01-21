package service

import (
	"errors"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/go-chi/chi"

	realize_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/logic"
)

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

//authenticationToken 验证令牌与UID
func AuthenticationToken(r *http.Request) error {
	//1.接收值
	t, _ := r.Cookie("token")
	u, _ := r.Cookie("uid")
	//2.处理值
	us := &realize_logic.User{}
	us.UserID = u.Value
	if err := us.VerifyTheUser(t.Value); err != nil {
		return err
	}
	return nil
}

// AdminLevel admin  权限等级验证
func AdminLevel(r *http.Request) error {
	user := &realize_logic.User{}
	lockID := chi.URLParam(r, "userID")
	u, _ := r.Cookie("uid")
	if lockID == u.Value {
		return errors.New("Operation account is equal to login account")
	}
	user.UserID = u.Value
	result, err := user.CurrentUserInformation()
	if err != nil {
		return err
	}
	logrus.Println(result.Role)
	if result.Role != "admin" {
		return errors.New("Role permissions error")
	}
	return nil
}

// ManagerLevel Manager  权限等级验证
func ManagerLevel(r *http.Request) error {
	user := &realize_logic.User{}
	lockID := chi.URLParam(r, "userID")
	u, _ := r.Cookie("uid")
	if lockID == u.Value {
		return errors.New("Operation account is equal to login account")
	}
	user.UserID = u.Value
	result, err := user.CurrentUserInformation()
	if err != nil {
		return err
	}
	logrus.Println(result.Role)
	if result.Role != "manager" {
		return errors.New("Role permissions error")
	}
	return nil
}

// EditorLevel Editor  权限等级验证
func EditorLevel(r *http.Request) error {
	user := &realize_logic.User{}
	lockID := chi.URLParam(r, "userID")
	u, _ := r.Cookie("uid")
	if lockID == u.Value {
		return errors.New("Operation account is equal to login account")
	}
	user.UserID = u.Value
	result, err := user.CurrentUserInformation()
	if err != nil {
		return err
	}
	logrus.Println(result.Role)
	if result.Role != "editor" {
		return errors.New("Role permissions error")
	}
	return nil
}
