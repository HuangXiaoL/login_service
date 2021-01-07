package with_user_structure

//Register 注册需要参数
type Register struct {
	Email    string
	Password string
}

type Account interface {
	RegisterUserInfo() // 注册用户信息
}
