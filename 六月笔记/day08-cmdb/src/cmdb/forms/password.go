package forms

import (
	"cmdb/models"
	"regexp"

	"github.com/astaxie/beego/validation"
)

/*
	验证密码的forms
*/

// 接收form表单页面传递过来的数据,form表单有几个这里就搞几个
// form 是在controller表单中引用的
type PasswordModifyForm struct {
	User         *models.User `form:"-"`
	OldPassword  string       `form:"old_password"`
	NewPassword1 string       `form:"new_password1"`
	NewPassword2 string       `form:"new_password2"`
}

// Valid 数据检查
func (f *PasswordModifyForm) Valid(validation *validation.Validation) {

	// 验证旧密码
	if ok := f.User.VaildPassword(f.OldPassword); !ok {
		validation.AddError("default,default,default", "旧密码错误")
		return
	}
	// 验证密码范围
	passwordRegex := "^[0-9a-zA-Z_.\\$\\!#%^\\*\\(\\)\\+]{6,20}$"
	validation.Match(f.NewPassword1, regexp.MustCompile(passwordRegex),
		"defalut.defalut.defalut").Message("密码格式不正确")
	if validation.HasErrors() {
		return
	} else if f.NewPassword1 != f.NewPassword2 {
		validation.AddError("default.dafault", "新旧密码不能一致")
	} else if f.OldPassword == f.NewPassword1 {
		// 验证新旧密码不能一致
		validation.AddError("default.dafault", "新旧密码不能一致")

	}
}
