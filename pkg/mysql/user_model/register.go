package user_model

import "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure"

var (
	_ with_user_structure.Account = (*User)(nil)
)

type User struct {
}

func (u *User) RegisterUserInfo() {

}
