package user_model

import (
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/model"
)

var (
	_ model.AccountInformation = (*UserInfo)(nil)
)

type UserInfo struct {
	model.UserInfo
}

func Get() *UserInfo {
	return &UserInfo{}
}

//CreateUserInfo 注册用户
func (u *UserInfo) CreateUserInfo() error {
	sqlStr := "INSERT INTO user_info(uuid, password,password_salt,email) VALUES(?, ?,?, ?)"
	_, err := db.Exec(sqlStr, u.UserID, u.Password, u.PasswordSalt, u.Email)
	if err != nil {
		return err
	}
	return err
}

//CreateUserLoginInfo 用户登录信息
func (u *UserInfo) CreateUserLoginInfoByEmail() (err error) {
	sqlStr := "UPDATE user_info SET session_salt = ? WHERE email = ?"
	_, err = db.Exec(sqlStr, u.SessionSalt, u.Email)
	if err != nil {
		return
	}
	return
}

//SelectUserInfoByEmail 根据email 查询用户信息
func (u *UserInfo) SelectUserInfoByEmail() (model.UserInfo, error) {
	sqlStr := "SELECT uuid, password, password_salt,email FROM user_info WHERE  email= ?"
	user := model.UserInfo{}
	if err := db.Get(&user, sqlStr, u.Email); err != nil {
		return model.UserInfo{}, err
	}
	return user, nil
}
