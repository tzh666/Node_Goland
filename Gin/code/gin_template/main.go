package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 默认路由
	r := gin.Default()

	// gin框架自定义模板函数
	r.SetFuncMap(template.FuncMap{
		// 自定义函数名 safe
		// 匿名函数 func (str string) template.HTML ---> 自定义模板函数实现的功能
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})

	// gin框架模板解析,Gin框架中使用LoadHTMLGlob()或者LoadHTMLFiles()方法进行HTML模板渲染。
	r.LoadHTMLFiles("template/index.html")

	r.GET("/index", func(c *gin.Context) {
		// HTTP请求
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "<a href='https://baidu.com'>tzh666</a>",
		})
	})

	r.Run(":9090")
}
