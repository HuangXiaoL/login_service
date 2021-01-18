package model

import (
	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"

	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/model"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user_model"
)

//Register 注册
func Register(u map[string]interface{}) error {
	//组装数据
	logrus.Println(u)
	uinfo := user_model.Get()
	uinfo.Password = typeJudgment(u["Password"])
	uinfo.PasswordSalt = typeJudgment(u["PasswordSalt"])
	uinfo.UserID = typeJudgment(u["UserID"])
	uinfo.Email = typeJudgment(u["Email"])
	//创建用户
	logrus.Printf("%+v", uinfo)
	if err := uinfo.CreateUserInfo(); err != nil {
		return err
	}
	return nil
}

//LoginSalt 设置登录盐 session salt
func LoginSalt(email string, salt string) (err error) {
	uinfo := user_model.Get()
	uinfo.Email = email
	uinfo.SessionSalt = salt
	if err = uinfo.CreateUserLoginInfoByEmail(); err != nil {
		return err
	}
	return
}

//FindUserBuyEmail 查询用户，根据用户邮箱
func FindUserBuyEmail(email string) (model.UserInfo, error) {
	uinfo := user_model.UserInfo{}
	uinfo.Email = email
	us, err := uinfo.SelectUserInfoByEmail()
	if err != nil { //不存在返回错误信息
		return model.UserInfo{}, err
	}
	return us, nil
}

//FindUserBuyEmail 查询用户，根据用户UID
func FindUserBuyUID(uid string) (model.UserInfo, error) {
	uinfo := user_model.UserInfo{}
	uinfo.UserID = uid
	logrus.Info(uinfo.UserID)
	us, err := uinfo.SelectUserInfoByUID()
	if err != nil { //不存在返回错误信息
		return model.UserInfo{}, err
	}
	return us, nil
}

//UpdateMyPassword 更新密码
func UpdateMyPassword(uid, newPassword, passwordSalt string) (err error) {
	uinfo := user_model.UserInfo{}
	uinfo.UserID = uid
	uinfo.Password = newPassword
	uinfo.PasswordSalt = passwordSalt
	if err = uinfo.UpdatePasswordAndPasswordSaltByUID(); err != nil {
		return
	}
	return
}

//typeJudgment 类型断言
func typeJudgment(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}
