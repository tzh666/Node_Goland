package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 获取浏览器的请求，获取ur发起请求携带的query string 参数
	r.GET("/web", func(c *gin.Context) {
		// 通过c.Query获取请求中携带的参数  <http://127.0.0.1:9090/web?query=tzh>
		// name := c.Query("query")

		// 取不到就用默认值  <http://127.0.0.1:9090/web?xxx>
		// name := c.DefaultQuery("query", "zhangsan")

		// GetQuery 有两个返回值,可做判断用 <http://127.0.0.1:9090/web?query=tzh>
		name, ok := c.GetQuery("query")
		if !ok {
			fmt.Println("取值失败")
			return
		}
		// <http://127.0.0.1:9090/web?query=tzh&age=18>
		age, _ := c.GetQuery("query")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})

	r.Run(":9090")
}
