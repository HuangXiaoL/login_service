package logic

import (
	"crypto/md5"
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/go-basic/uuid"
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/logic"
)

//User 注册用户所需结构体
type User struct {
	UserID       string //用户ID
	Password     string //用户密码
	PasswordSalt string //用户密码盐
	Email        string //用户邮箱
}

var (
	_ logic.UserBehavior = (*User)(nil)
)

//RegisterInfo 注册用户数据出来
func (u *User) RegisterInfo(r logic.Register) (err error) {
	u.UserID = uuid.New()
	u.Email = r.Email
	u.PasswordSalt = uuid.New()
	u.Password = u.passwordSaltDispose(r.Password)
	logrus.Printf("%+v", u)
	return
}

//passwordSaltDispose 密码加盐 p 密码原始字符串
func (u *User) passwordSaltDispose(p string) string {
	data := []byte(p + u.PasswordSalt)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
