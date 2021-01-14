package logic

import (
	"crypto/md5"
	"fmt"
	"reflect"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/jwt"

	"github.com/go-basic/uuid"
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/model"
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/logic"
)

var (
	_ logic.UserBehavior = (*User)(nil)
)

//User 注册用户所需结构体
type User struct {
	UserID       string //用户ID
	PasswordSalt string //用户密码盐
	Email        string `valid:"email,optional"` //邮箱
	Password     string `valid:"optional"`       //密码
}

//RegisterInfo 注册用户数据处理
func (u *User) RegisterInfo(r logic.Register) (err error) {
	//logrus.Printf("%+v", r)
	u.UserID = uuid.New()
	u.Email = r.Email
	u.PasswordSalt = uuid.New()
	u.Password = u.passwordSaltDispose(r.Password)
	userInfo := u.structToMap()                     //结构体转map
	if err = model.Register(userInfo); err != nil { // 注册用户
		return err
	}
	return
}

//Login 登录数据处理 logic.Login 登录所需参数结构体， way 登录方式 1 注册后不需要效验，执行登录
func (u *User) Login(login logic.Login, way int) (s string, err error) {
	sSalt := uuid.New()
	if way == 1 { //注册后登录
		s, err = jwt.GenToken(u.UserID, []byte(sSalt), 1)
		if err != nil {
			return
		}
		if err = model.LoginSalt(login.Email, sSalt); err != nil {
			return
		}
		return
	}
	//正常登录流程
	//jwt.GenToken()
	return
}

//passwordSaltDispose 密码加盐 p 密码原始字符串
func (u *User) passwordSaltDispose(p string) string {
	data := []byte(p + u.PasswordSalt)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

//structToMap 结构体转 map
func (u User) structToMap() map[string]interface{} {
	obj1 := reflect.TypeOf(u)
	obj2 := reflect.ValueOf(u)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
