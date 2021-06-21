package forms

// 因为提交的URL是小写得打标签,form:name,中间不能有空格
// 用户登录表单
type LoginForm struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}
