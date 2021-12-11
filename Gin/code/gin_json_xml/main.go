package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// gin.H 是map[string]interface{}的缩写
	r.GET("/json1", func(c *gin.Context) {
		// 方式一,自己拼接JSON格式数据
		data := map[string]interface{}{
			"name": "tzh",
			"age":  18,
		}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/json2", func(c *gin.Context) {
		// 方式一,自己拼接JSON格式数据
		c.JSON(http.StatusOK, gin.H{
			"name": "tzh1",
			"age":  18,
		})
	})
	// 方式二：使用结构体,首字母大写才能传到前端去,但是可以用 `` 做灵活定制化操作,这样返回给前端的就是name  password。
	type msg struct {
		Name     string `json:"name" xml:"name"`
		Password string `json:"passwoed" xml:"password"`
	}
	r.GET("/json3", func(c *gin.Context) {
		// 方式一,自己拼接JSON格式数据
		data := msg{
			"tzh",
			"tzh",
		}
		// JSON序列号
		c.JSON(http.StatusOK, data)
	})

	r.Run(":9090")
}
