package model

import (
	"encoding/json"
	"strconv"

	"gitlab.haochang.tv/huangxiaolei/login_service/pkg/mysql/user"

	"github.com/sirupsen/logrus"

	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/model"
)

//Register 注册
func Register(u map[string]interface{}) error {
	//组装数据
	logrus.Println(u)
	uinfo := user.Get()
	uinfo.Password = typeJudgment(u["Password"])
	uinfo.PasswordSalt = typeJudgment(u["PasswordSalt"])
	uinfo.UserID = typeJudgment(u["UserID"])
	uinfo.Email = typeJudgment(u["Email"])
	//创建用户
	return uinfo.CreateUserInfo()
}

//LoginSalt 设置登录盐 session salt
func LoginSalt(email string, salt string) (err error) {
	uinfo := user.Get()
	uinfo.Email = email
	uinfo.SessionSalt = salt
	if err = uinfo.CreateUserLoginInfoByEmail(); err != nil {
		return err
	}
	return
}

//FindUserBuyEmail 查询用户，根据用户邮箱
func FindUserBuyEmail(email string) (model.UserInfo, error) {
	uinfo := user.Get()
	uinfo.Email = email
	us, err := uinfo.SelectUserInfoByEmail()
	if err != nil { //不存在返回错误信息
		return model.UserInfo{}, err
	}
	return us, nil
}

//FindUserBuyUID 查询用户，根据用户UID
func FindUserBuyUID(uid string) (model.UserInfo, error) {
	uinfo := user.Get()
	uinfo.UserID = uid
	us, err := uinfo.SelectUserInfoByUID()
	if err != nil { //不存在返回错误信息
		return model.UserInfo{}, err
	}
	return us, nil
}

//FindUserInfoAndRoleBuyUID 查询用户信息和角色根据用户UID
func FindUserInfoAndRoleBuyUID(uid string) (model.UserInfo, model.Role, error) {
	uinfo := user.Get()
	uinfo.UserID = uid
	info, err := uinfo.SelectUserInfoByUID()
	if err != nil {
		return model.UserInfo{}, model.Role{}, err
	}
	uinfo.Role = info.Role
	role, err := uinfo.SelectRoleByID()
	if err != nil {
		return model.UserInfo{}, model.Role{}, err
	}
	return info, role, nil

}

//UpdateMyPassword 更新密码
func UpdateMyPassword(uid, newPassword, passwordSalt string) (err error) {
	uinfo := user.Get()
	uinfo.UserID = uid
	uinfo.Password = newPassword
	uinfo.PasswordSalt = passwordSalt
	if err = uinfo.UpdatePasswordAndPasswordSaltByUID(); err != nil {
		return
	}
	return
}

//UpdateUerLockTimeByUID 根据用户ID更新用户锁定时间
func UpdateUerLockTimeByUID(nowTime, account string) (err error) {
	uinfo := user.Get()
	uinfo.LockTime.String = nowTime                       // 锁定时间
	uinfo.UserID = account                                // 锁定账号
	if err = uinfo.UpdateUerLockTimeByUID(); err != nil { // 更新数据
		return
	}
	return
}

// FindRoleBuyRoleName 根据角色名称查询角色
func FindRoleBuyRoleName(name string) (role model.Role, err error) {
	uinfo := user.Get()
	role, err = uinfo.SelectRoleByRoleName(name)
	if err != nil {
		return
	}
	return
}

//UpdateUserRole 更新用户的角色
func UpdateUserRole(uid, roleID string) (err error) {
	uinfo := user.Get()
	uinfo.UserID = uid
	uinfo.Role = roleID
	if err = uinfo.UpdateUerRoleID(); err != nil {
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
