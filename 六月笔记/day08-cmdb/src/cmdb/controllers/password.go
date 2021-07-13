package controllers

import (
	"cmdb/base/controllers/auth"
	"cmdb/base/errors"
	"cmdb/forms"
	"cmdb/services"
	"html/template"

	"github.com/astaxie/beego/validation"
)

/*
	验证密码的controller
*/

// 用户修改密码控制器
type PasswordController struct {
	auth.LayouController
}

// 修改用户密码
func (c *PasswordController) Modify() {
	form := &forms.PasswordModifyForm{User: c.LoginUser}
	errs := errors.New()
	text := ""
	// 同样的Get请求加载页面,Post请求更新数据
	if c.Ctx.Input.IsPost() {
		if err := c.ParseForm(form); err == nil {
			valid := &validation.Validation{}
			if hasErrors, err := valid.Valid(form); err != nil {
				errs.Add("default", err.Error())
			} else if !hasErrors {
				errs.AddAalidation(valid)
			} else {
				services.UserService.ModifyPassword(c.LoginUser.ID, form.NewPassword1)
				text = "修改密码成功"
			}
		}
	}
	c.TplName = "password/modify.html"
	// 把错误信息回显到页面
	c.Data["errors"] = errs
	c.Data["text"] = text

	// 生成一个input标签 <input type="hidden" name="_xsrf" value="QGJlODKVegKMsRcJtToDt6ppPVCYSL9Z" />
	c.Data["xsrf_input"] = template.HTML(c.XSRFFormHTML())
	// fmt.Println(template.HTML(c.XSRFFormHTML()))
	c.Data["title"] = "修改用户密码"

	// 自定义样式,script脚本。在views/password目录下
	c.LayoutSections["SectionStyle"] = "password/modify_style.html"
	c.LayoutSections["SectionScript"] = "password/modify_script.html"
}
