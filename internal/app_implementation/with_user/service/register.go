package service

import (
	"net/http"

	realize_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_implementation/with_user/logic"

	"github.com/sirupsen/logrus"

	struct_logic "gitlab.haochang.tv/huangxiaolei/login_service/internal/app_structure/with_user_structure/logic"

	"github.com/joyparty/httpkit"
)

//RegisterUserInfo 注册用户
func RegisterUserInfo(w http.ResponseWriter, r *http.Request) {
	//1.上传参数赋值
	req := struct_logic.Register{}
	_ = r.ParseForm()
	if err := httpkit.ScanValues(&req, r.PostForm); err != nil { //绑定参数
		httpkit.WrapError(err).WithHeader("err", "Incorrect parameter, missing parameter, parameter content, or type").WithStatus(http.StatusBadRequest).Panic()
	}
	logrus.Printf("%+v", req)
	//2.数据处理
	//执行注册流程
	u := realize_logic.User{}
	if err := u.RegisterInfo(req); err != nil { //注册失败，uuid或者email重复 ，证明该邮箱已经被注册了，下行错误代码 409
		logrus.Println(err)
		httpkit.WrapError(err).WithStatus(http.StatusConflict).Panic()
	}
	//执行登录流程,生成session salt，下行201
	l := struct_logic.Login{}
	l.Email = u.Email
	u.Login(l, 1) //way 参数为1 不需要效验账号密码，直接登录
	//3.结果下行
}
