package routers

import (
	"cmdb/controllers"

	"github.com/astaxie/beego"
)

func init() {

	// 错误页面处理,访问错误页面的时候就会跳转到自定义error页面
	beego.ErrorController(&controllers.ErrorController{})

	// 访问 / (根路径)路径时,跳转到首页
	beego.Router("/", &controllers.HomeController{}, "*:Index")

	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.HomeController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PasswordController{})
}
