package main

import "github.com/astaxie/beego"

func main() {
	// 设置到文件中
	beego.SetLogger("file", `{"filename" : "test.log"}`)
	// 设置日志级别，只有比这个级别高的才会打印
	beego.SetLevel(beego.LevelInformational)
	// 会提示那一行的日志
	beego.SetLogFuncCall(true)
	beego.Debug("我是一个调试日志")
	beego.Informational("我是一个提醒日志")
	beego.Error("我是一个错误日志")
}
