package user_model

import (
	"gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/model"
)

var (
	_ model.AccountInformation = (*User)(nil)
)

type User struct {
}

func (u *User) RegisterUserInfo() {

}
