package controllers

import "cmdb/base/controllers/base"

// ErrorController 页面错误处理
type ErrorController struct {
	base.BaseController
}

// 404页面处理
func (c *ErrorController) Error404() {
	c.TplName = "error/404.html"
}

// 自定义跳转到错误页面处理,名字跟Abort中的一致
/*
	c.Abort("NotPermission")
	return
*/
func (c *ErrorController) ErrorNotPermission() {
	c.TplName = "error/not_permission.html"
}
