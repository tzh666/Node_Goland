package forms

// 因为提交的URL是小写得打标签,form:name,中间不能有空格
// 用户修改表单,现在定义了name,只能修改name的值
type UserModifyForm struct {
	ID         int    `form:"id"`
	StaffID    string `form:"staffid"`
	Name       string `form:"name"`
	Nickname   string `form:"nickname"`
	Password   string `form:"password"`
	Gender     int    `form:"gender"`
	Tel        string `form:"tel"`
	Addr       string `form:"addr"`
	Email      string `form:"email"`
	Department string `form:"department"`
}
