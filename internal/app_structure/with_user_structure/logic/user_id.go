package logic

//Register 注册需要参数
type Register struct {
	Email    string `valid:"email,optional"`
	Password string `valid:"optional"`
}

type UserBehavior interface {
	RegisterInfo(Register) (err error) //注册信息
}
