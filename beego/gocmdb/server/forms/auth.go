package forms

import (
	"strings"

	"github.com/imsilence/gocmdb/server/models"

	"github.com/astaxie/beego/validation"
)

type LoginForm struct {
	Name     string `form:"name"`     // 跟登录表单里面的name="namne" 要保持
	Password string `form:"password"` //

	User *models.User // 嵌套个User用于接收数据
}

// beego密码认证
func (f *LoginForm) Valid(v *validation.Validation) {
	// 去除前后空格
	f.Name = strings.TrimSpace(f.Name)
	f.Password = strings.TrimSpace(f.Password)
	// 判断是否为空
	if f.Name == "" || f.Password == "" {
		v.SetError("error", "用户名或者密码错误")
	} else {
		// 通过Name查找,ValidPasswprd
		if user := models.DefaultUserManager.GetByName(f.Name); user == nil || !user.ValidPassword(f.Password) {
			v.SetError("error", "用户名或者密码错误")
		} else if user.IsLock() {
			v.SetError("error", "用户已经被锁定")
		} else {
			f.User = user
		}
	}
}
