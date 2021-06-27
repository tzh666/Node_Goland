package auth

import (
	"fmt"

	"github.com/imsilence/gocmdb/server/controllers/base"
	"github.com/imsilence/gocmdb/server/models"
)

// 给其他controller判断是否登录用的
type LoginRequiredController struct {
	base.BaseController
	User *models.User
}

func (c *LoginRequiredController) Prepare() {
	// 重写Prepare
	c.BaseController.Prepare()

	if user := DefaultManger.IsLogin(c); user == nil {
		// 如果返回的用户信息不为空,未登录
		fmt.Println("用户信息不为空,未登录")
		DefaultManger.GoToLoginPage(c)
		c.StopRun()
	} else {
		// 否则是登录
		c.User = user
		c.Data["user"] = user
	}
}

// 登录验证 AuthController
type AuthController struct {
	base.BaseController
}

func (c *AuthController) Login() {
	// 判断是否登录ok
	DefaultManger.Login(c)

}

func (c *AuthController) Logout() {
	DefaultManger.Logout(c)
}
