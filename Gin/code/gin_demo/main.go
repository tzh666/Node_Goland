package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello golang",
	})
}

func main() {
	// 返回一个默认的路由引擎
	r := gin.Default()

	// 使用GET方法
	r.GET("/hello", sayHello)

	// 约定用RESTful API做接口开发, 使用POSTMAN去调用我们的接口,用不同的方式请求,返回不同的信息
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"method": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "POST",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"method": "PUT",
		})
	})

	// 启动函数,默认9090
	// r.Run("8888") 自定义端口
	r.Run()
}
