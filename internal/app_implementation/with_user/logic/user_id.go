package logic

import (
	"crypto/md5"
	"fmt"
	"reflect"

	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/model"

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

//RegisterInfo 注册用户数据出来
func RegisterInfo(r logic.Register) (err error) {
	u := User{}
	u.UserID = uuid.New()
	u.Email = r.Email
	u.PasswordSalt = uuid.New()
	u.Password = u.passwordSaltDispose(r.Password)
	userInfo := u.structToMapDemo()
	if err := model.UserInfo(userInfo); err != nil {
		return err
	}
	return
}

//passwordSaltDispose 密码加盐 p 密码原始字符串
func (u *User) passwordSaltDispose(p string) string {
	data := []byte(p + u.PasswordSalt)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func (u User) structToMapDemo() map[string]interface{} {
	obj1 := reflect.TypeOf(u)
	obj2 := reflect.ValueOf(u)
	var data = make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		data[obj1.Field(i).Name] = obj2.Field(i).Interface()
	}
	return data
}
