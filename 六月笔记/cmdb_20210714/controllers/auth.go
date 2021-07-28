package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/astaxie/beego"

	"cmdb/base/controllers/base"
	"cmdb/base/errors"
	"cmdb/config"
	"cmdb/forms"
	"cmdb/services"
)

// 用户登录认证相关的controller
type AuthController struct {
	base.BaseController
}

/*
	思路：首先判断是什么请求
	1、GET请求,加载页面
	2、POST请求，做认证------> 判断如果是POST请求---->c.Ctx.Input.IsPost()
	3、获取用户提交的数据
		 a、认证成功怎么样
		 b、认证失败怎么样
*/
// Login 登录认证
func (c *AuthController) Login() {

	// 如果用户已经登录,则直接跳转到用户列表页面
	if sessionUser := c.GetSession(beego.AppConfig.DefaultString("auth::SessionKey", "user")); sessionUser != nil {
		c.Redirect(beego.URLFor("UserController.Query"), 302)
	}
	// 定义成指针
	form := &forms.LoginForm{}
	errs := errors.New()
	if c.Ctx.Input.IsPost() {

		// 缓存配置,每次登录都+1,默认0
		config.Cache.Incr("login")
		fmt.Println(config.Cache.Get("login"))

		// 获取用户提交的数据,如果读取数据发生错误,则显示错误信息
		if err := c.ParseForm(form); err == nil {
			// 调用函数获取用户传递进来的name，赋值给user
			user := services.UserService.GetByName(form.Name)
			if user == nil {
				errs.Add("default", "用户名或者密码错误")
				// 日志打印
				beego.Error(fmt.Sprintf("用户不存在: %s", form.Name))

				// 用户不存在
			} else if user.VaildPassword(form.Password) {
				beego.Informational(fmt.Sprintf("用户认证成功: %s", form.Name))
				// 记录用户状态(session 记录在服务器端的)
				c.SetSession("user", user.ID) // 把User结构体的ID属性存到session中

				// 密码正确，跳转到首页/home   302
				c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
			} else {
				errs.Add("default", "用户名或者密码错误")
				// 密码不正确
				beego.Error(fmt.Sprintf("用户名或者密码错误: %s", form.Name))
			}
		} else {
			errs.Add("default", "用户名或者密码错误")
		}
	}

	// c.Data[] 渲染数据到模板上
	c.Data["form"] = form
	c.Data["errors"] = errs
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	// GET请求直接加载页面,指定要加载的页面,记得默认去views问价夹下寻找,文件名前面有一层目录的话也要加进去-----> views/xxx/test.html
	// 定义加载页面
	c.TplName = "auth/login.html"
}

// 退出登录
func (c *AuthController) Logout() {
	// 退出登录,销毁session,DestroySession删除所有session
	c.DestroySession()
	// 跳转到登录页面
	c.Redirect(beego.URLFor("AuthController.Login"), http.StatusFound)
}
