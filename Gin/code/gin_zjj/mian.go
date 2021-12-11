package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 平时写这个常用的也是一个 gin.HandlerFunc类型  <也是一个中间件>
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"x": "1",
	})
}

// 自定义中间件,统计耗时
func m1(c *gin.Context) {
	start := time.Now()

	// 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
	c.Set("name", "小王子")

	// 调用该请求的剩余处理程序
	c.Next()

	// 不调用该请求的剩余处理程序
	// c.Abort()
	// 计算耗时
	cost := time.Since(start)
	fmt.Println(cost)
}

func main() {
	r := gin.Default()

	// 注册一个全局中间件 m1
	r.Use(m1)

	// GET(relativePath string, handlers ...HandlerFunc)
	r.GET("/index", m1, indexHandler) // <会先执行m1>

	r.Run(":9090")
}
