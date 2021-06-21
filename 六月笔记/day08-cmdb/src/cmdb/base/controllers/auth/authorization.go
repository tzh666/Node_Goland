package auth

import (
	"cmdb/base/controllers/base"
	"cmdb/models"
	"cmdb/services"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
)

// AuthorizationController 所有需要认证才能访问的基础控制器
type AuthorizationController struct {
	base.BaseController
	LoginUser *models.User
}

// 公共的地方就放在公共的地方
func (c *AuthorizationController) getNav() string {

	// 返回值第一个是控制器名称controllerName，第二个是调用的函数名称actionName
	controllerName, _ := c.GetControllerAndAction()

	// 把Controller去掉,(controllerName==UserController),保留User
	return strings.ToLower(strings.TrimPrefix(controllerName, "Controller"))
}

// Prepare 用户登录session认证检查
func (c *AuthorizationController) Prepare() {

	// 读取配置文件的值，默认是user
	sessionKey := beego.AppConfig.DefaultString("auth::SessionKey", "user")
	sessionValue := c.GetSession(sessionKey)

	// 初始化
	c.Data["loginUser"] = nil
	c.Data["nav"] = c.getNav()

	// 如果sessionValue 不等于nil,说明登录ok
	if sessionValue != nil {
		// 登录成功再断言赋值给pk
		if pk, ok := sessionValue.(int); ok {
			// 再判断用户是否为nil
			if user := services.UserService.GetByPk(pk); user != nil {
				// 通过ID查询到user,再通过模板传参到页面
				c.Data["loginUser"] = user
				c.LoginUser = user
				// user != nil,则return
				return
			}
		}
	}
	// 不符合上面的判断,说明没登录成功
	action := beego.AppConfig.DefaultString("auth::LoginAction",
		"AuthController.Login")
	c.Redirect(beego.URLFor(action), http.StatusFound)
}
