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
	if err := realize_logic.RegisterInfo(req); err != nil { //查到了数据 ，证明该邮箱已经被注册了，下行错误代码 409
		httpkit.WrapError(err).WithStatus(http.StatusConflict).Panic()
	}

	//3.结果下行
}
