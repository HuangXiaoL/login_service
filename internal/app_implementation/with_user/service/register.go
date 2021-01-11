package service

import (
	"net/http"

	realize_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/logic"

	struct_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/logic"

	"github.com/joyparty/httpkit"
	"github.com/sirupsen/logrus"
)

//RegisterUserInfo 注册用户
func RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	//1.上传参数赋值
	req := struct_logic.Register{}
	_ = r.ParseForm()
	if err := httpkit.ScanValues(&req, r.PostForm); err != nil { //绑定参数
		logrus.Println(err)
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	logrus.Printf("%+v", req)
	//2.数据处理
	u := &realize_logic.User{}
	u.RegisterInfo(req)

	//3.结果下行
}
