package model

import "database/sql"

//UserInfo 用户表信息
type UserInfo struct {
	UserID       string         `db:"uuid"`          //用户ID
	Password     string         `db:"password"`      //用户密码
	PasswordSalt string         `db:"password_salt"` //用户密码盐
	SessionSalt  string         `db:"session_salt"`  //用户密码盐
	Email        string         `db:"email"`         //用户邮箱
	CreateTime   string         `db:"create_time"`   // 创建时间
	Role         string         `db:"role"`          //角色ID
	LockTime     sql.NullString `db:"lock_time"`     // 锁定时间
}

//Role 角色表
type Role struct {
	RoleID   string `db:"id"`        //角色名称
	RoleName string `db:"role_name"` //角色名称
	RoleAuth string `db:"role_auth"` //角色权限
}

//AccountInformation 账号信息操作相关接口
type AccountInformation interface {
	CreateUserInfo() error                   // 注册用户信息
	CreateUserLoginInfoByEmail() (err error) //登录状态 session salt 创建

	SelectUserInfoByEmail() (UserInfo, error) //查询用户信息根据email
	SelectUserInfoByUID() (UserInfo, error)   //查询用户信息根据uuid
	SelectRoleByID() (Role, error)            //查询角色根据角色ID

	UpdatePasswordAndPasswordSaltByUID() (err error) // 修改密码和密码盐根据uuid
	UpdateUerLockTimeByUID() (err error)             // 锁定账号
}
