package logic

//Register 注册需要参数
type Register struct {
	Email    string `valid:"email,optional"`
	Password string `valid:"optional"`
}

//Login 登录所需参数
type Login struct {
	Email    string `valid:"email,optional"`
	Password string `valid:"optional"`
	Remember int    `valid:""`
}
type ChangePassword struct {
	Password string `json:"old_password" valid:"optional"`
}
type UserBehavior interface {
	RegisterInfo(Register) (err error)      //注册信息
	Login(Login, int) (string, error)       //登录，生成session salt 生成json web token
	VerifyTheUser(token string) (err error) //验证用户token和UID的正确性，确保token 和UID 匹配
	MyPassword(newPWD string)               //处理用户自己的密码
}
