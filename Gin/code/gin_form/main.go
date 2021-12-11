package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.POST("/Login", func(c *gin.Context) {
		username := c.PostForm("username")
		address := c.PostForm("address")

		// DefaultPostForm取不到值时会返回指定的默认值
		username1 := c.DefaultPostForm("username", "森森")
		address1 := c.DefaultPostForm("address", "北京")

		username2, _ := c.GetPostForm("username")
		address2, _ := c.GetPostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":   "ok",
			"Nsername":  username,
			"Address":   address,
			"Nsername1": username1,
			"Address1":  address1,
			"Nsername2": username2,
			"Address2":  address2,
		})
	})

	r.POST("/json", func(c *gin.Context) {
		// 注意：下面为了举例子方便，暂时忽略了错误处理
		b, _ := c.GetRawData() // 从c.Request.Body读取请求数据
		// 定义map或结构体
		var m map[string]interface{}
		// 反序列化
		_ = json.Unmarshal(b, &m)

		c.JSON(http.StatusOK, m)
	})
	r.Run(":8080")
}
