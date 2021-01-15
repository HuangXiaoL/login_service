package model

type UserInfo struct {
	UserID       string `db:"uuid"`          //用户ID
	Password     string `db:"password"`      //用户密码
	PasswordSalt string `db:"password_salt"` //用户密码盐
	SessionSalt  string `db:"session_salt"`  //用户密码盐
	Email        string `db:"email"`         //用户邮箱
}

//Account 账号信息操作相关接口
type AccountInformation interface {
	CreateUserInfo() error                    // 注册用户信息
	CreateUserLoginInfoByEmail() (err error)  //登录状态 session salt 创建
	SelectUserInfoByEmail() (UserInfo, error) //查询用户信息根据email
	SelectUserInfoByUID() (UserInfo, error)   //查询用户信息根据uuid
}
