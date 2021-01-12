package model

import (
	"encoding/json"
	"strconv"

	"github.com/sirupsen/logrus"
	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user_model"
)

func UserInfo(u map[string]interface{}) error {
	//组装数据
	//logrus.Println(u)
	uinfo := user_model.UserInfo{}
	uinfo.Password = TypeJudgment(u["Password"])
	uinfo.PasswordSalt = TypeJudgment(u["PasswordSalt"])
	uinfo.UserID = TypeJudgment(u["UserID"])
	uinfo.Email = TypeJudgment(u["Email"])
	//logrus.Printf("%+v", uinfo)

	//查询用户是否存在
	us, err := uinfo.SelectUserInfoByEmail()
	if err == nil { //用户存在返回错误信息改用户存在
		return err

	}
	logrus.Printf("%+v", us)
	//创建用户
	//if err := uinfo.CreateUserInfo(); err != nil {
	//	return err
	//}
	return nil
}
func TypeJudgment(value interface{}) string {
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
