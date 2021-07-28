package controllers

import "github.com/astaxie/beego"

type ApiController struct {
	beego.Controller
}

func (c *ApiController) Prepare() {
	// 可以在这里验证token

}
