package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/forms"
	"cmdb/services"
	"net/http"

	"github.com/astaxie/beego"
)

// UserController 用户管理控制器
type UserController struct {
	auth.LayouController
}

// 新增用户
func (c *UserController) Add() {
	form := &forms.UserModifyForm{}
	c.Data["formadd"] = form
	c.TplName = "user/add.html"
}

// 根据ID删除用户
func (c *UserController) Delete() {
	// 获取url上的pk,进行删除
	// 如果删除的是当前登录用户则,删除失败
	if pk, err := c.GetInt("pk"); err == nil && c.LoginUser.ID != pk {
		// controller--->到model层删除数据了
		services.UserService.Delete(pk)
	}
	c.Redirect(beego.URLFor("UserController.Query"), http.StatusFound)
}

// 查询用户
func (c *UserController) Query() {

	/*
		读取flash存储的cookie值,c.Controller也是beego.controller
		flash := beego.ReadFromRequest(&c.Controller)
		flash.Data类型是一个切片map,可以直接在页面读取  ----> query.html
		fmt.Println("flash", flash.Data)
	*/

	beego.ReadFromRequest(&c.Controller)

	// 从页面提交的url中获取q的值，并且赋值给q，然后在请求的时候判断怎么查询即可
	q := c.GetString("q")
	c.Data["users"] = services.UserService.Query(q) // 查询所有数据
	c.Data["q"] = q                                 // 通过字段查询的数据

	c.TplName = "user/query.html"
}

// 编辑用户
func (c *UserController) Modify() {
	// 通过 定义的结构体映射model的字段（数据库的字段）修改,就是不直接修改数据库
	form := &forms.UserModifyForm{}
	// Get 获取用户数据
	// Post 修改数据
	if c.Ctx.Input.IsPost() { // 如果是Post请求,就直接修改
		if err := c.ParseForm(form); err == nil {
			// 验证数据,然后update数据
			services.UserService.Modify(form)

			// 存储消息,修改成功后的给用户一个提示
			// 这个消息存在cookie中,Set-Cookie: BEEGO_FLASH=%00notice%23BEEGOFLASH%23%E4%BF%AE%E6%94%B9%E7%94%A8%E6%88%B7%E4%BF%A1%E6%81%AF%E6%88%90%E5%8A%9F%00; Path=/
			flash := beego.NewFlash()
			// key value 形式
			flash.Set("notice", "修改用户信息成功")
			// c.Controller是beego的Controller
			flash.Store(&c.Controller)

			// 修改成功后302
			c.Redirect(beego.URLFor("UserController.Query"), 302)
		}

	} else if pk, err := c.GetInt("pk"); err == nil { // Get 请求就先展示修改页面
		if user := services.UserService.GetByPk(pk); user != nil { // 通过ID查询到用户信息
			// 页面显示
			form.ID = user.ID
			form.Name = user.Name
			form.Tel = user.Tel
			form.Addr = user.Addr
			form.Email = user.Email
		}
	}
	c.Data["form"] = form
	c.Data["xsrf_token"] = c.XSRFToken()
	c.TplName = "user/modify.html"
}
