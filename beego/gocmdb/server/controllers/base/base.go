package base

import (
	"html/template"

	"github.com/astaxie/beego"
)

// 基础Controller,组合了beego的控制器
type BaseController struct {
	beego.Controller
}

// csrf 认证,配置一个全局的认证,Post请求就会自动认证
func (c *BaseController) Prepare() {
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
}
