package main

import (
	"fmt"

	"github.com/astaxie/beego/validation"
)

func main() {
	valid := &validation.Validation{}

	test := ""
	// 验证test是否跟"a.c.v"相等, 用 . 分割
	valid.Required(test, "a.c.v").Key("Required").Message("输入内容不能为空")
	// 验证是否是英文字母
	valid.Alpha("1123", "alpha.alpha.alpha").Message("输入的内容只能是英文字母大小写跟数字")
	// 数字
	valid.Tel("1123", "tel.tel.tel").Message("输入的内容只能是大小写跟数字")

	// 获取验证结果
	if valid.HasErrors() {
		fmt.Println(valid.Errors)
		fmt.Println(valid.ErrorsMap)
	}

}
