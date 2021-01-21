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

//ChangePassword 修改密码所需参数
type ChangePassword struct {
	Password string `json:"new_password" valid:"optional"`
}

//SetRole 设置角色所需参数
type SetRole struct {
	Role string `json:"role" valid:"optional"`
}

//UserInfo 用户信息，查询用户信息下行数据
type UserInfo struct {
	ID       string `json:"id"`        //用户ID
	Email    string `json:"email"`     //邮箱
	CreateAt int64  `json:"create_at"` // 创建时间
	Role     string `json:"role"`      //权限角色名称
}

//UserBehavior 用户操作行为
type UserBehavior interface {
	RegisterInfo(Register) (err error)           //注册信息
	Login(Login, int) (string, error)            //登录，生成session salt 生成json web token
	VerifyTheUser(token string) (err error)      //验证用户token和UID的正确性，确保token 和UID 匹配
	MyPassword(newPWD string) (err error)        //处理用户自己的密码
	CurrentUserInformation() (UserInfo, error)   //获取该登录账号的信息，包括--用户ID，邮箱，创建时间，权限角色名称
	LockTheAccount(account string) (err error)   //锁定账号 传入参数  需要锁定的账号 account
	UNLockTheAccount(account string) (err error) //解锁账号 传入参数  需要解锁的账号 account
	SetUserRole(uid string, name string) error   // 设置账户角色
	SetDefaultPassword(uid string)               //设置默认密码
}
