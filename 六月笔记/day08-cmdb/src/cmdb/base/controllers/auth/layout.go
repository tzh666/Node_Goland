package auth

import "github.com/astaxie/beego"

// LayouController 布局控制器
type LayouController struct {
	// 肯定是登录以后才能看到的界面,所以得先认证
	AuthorizationController
}

// Prepare
func (c *LayouController) Prepare() {
	// 先做认证
	c.AuthorizationController.Prepare()
	c.Layout = "base/layouts/layout.html"

	// 自定义样式,script脚本。在views/password目录下
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["SectionStyle"] = ""
	c.LayoutSections["SectionScript"] = ""

	// 定义默认title
	c.Data["title"] = beego.AppConfig.DefaultString("AppName", "CMDB")
}
