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
type UserBehavior interface {
	RegisterInfo(Register) (err error) //注册信息
	Login(Login, int) (string, error)  //登录，生成session salt 生成json web token
}
