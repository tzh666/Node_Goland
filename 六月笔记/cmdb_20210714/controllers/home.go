package controllers

import (
	"cmdb/base/controllers/auth"
)

// 首页控制器
type HomeController struct {
	auth.LayouController
}

// 传递nva参数,页面跳转判断用
// func (c *HomeController) Prepare() {
// 	c.AuthorizationController.Prepare()
// 	c.Data["nav"] = "home"
// }

func (c *HomeController) Index() {
	c.TplName = "home/index.html"
}
