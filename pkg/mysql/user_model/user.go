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

//Get 获取结构体
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
	sqlStr := "SELECT uuid, password, password_salt,email,lock_time FROM user_info WHERE  email= ?"
	user := model.UserInfo{}
	if err := db.Get(&user, sqlStr, u.Email); err != nil {
		return model.UserInfo{}, err
	}
	return user, nil
}

//SelectUserInfoByUID 根据UID 查询用户信息
func (u *UserInfo) SelectUserInfoByUID() (model.UserInfo, error) {
	sqlStr := "SELECT uuid, password, password_salt,session_salt,email,role,create_time FROM user_info WHERE  uuid= ?"
	user := model.UserInfo{}
	if err := db.Get(&user, sqlStr, u.UserID); err != nil {
		return model.UserInfo{}, err
	}
	return user, nil
}

//UpdatePasswordAndPasswordSaltByUID 修改密码和密码盐
func (u *UserInfo) UpdatePasswordAndPasswordSaltByUID() (err error) {
	sqlStr := "UPDATE user_info SET password = ?,password_salt = ? WHERE uuid = ?"
	_, err = db.Exec(sqlStr, u.Password, u.PasswordSalt, u.UserID)
	if err != nil {
		return
	}
	return
}

//SelectRoleByID 根据ID 查询角色
func (u *UserInfo) SelectRoleByID() (model.Role, error) {
	sqlStr := "SELECT id,role_name,role_auth FROM role WHERE  id= ?"
	role := model.Role{}
	if err := db.Get(&role, sqlStr, u.Role); err != nil {
		return model.Role{}, err
	}
	return role, nil
}

//SelectRoleByRoleName 根据  名称  查询角色
func (u *UserInfo) SelectRoleByRoleName(name string) (model.Role, error) {
	sqlStr := "SELECT  id,role_name,role_auth FROM role WHERE  role_name = ?"
	role := model.Role{}
	if err := db.Get(&role, sqlStr, name); err != nil {
		return model.Role{}, err
	}
	return role, nil
}

//UpdateUerLockTimeByUID 锁定角色 更新锁定时间
func (u *UserInfo) UpdateUerLockTimeByUID() (err error) {
	sqlStr := "UPDATE user_info SET lock_time = ? WHERE uuid = ?"
	_, err = db.Exec(sqlStr, u.LockTime, u.UserID)
	if err != nil {
		return
	}
	return
}

//UpdateUerRoleID 更新用户的角色
func (u *UserInfo) UpdateUerRoleID() (err error) {
	sqlStr := "UPDATE user_info SET role = ? WHERE uuid = ?"
	_, err = db.Exec(sqlStr, u.Role, u.UserID)
	if err != nil {
		return
	}
	return
}
