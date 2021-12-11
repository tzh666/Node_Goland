package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("views/index.html")
	r.GET("/uploadfile", func(c *gin.Context) {
		// 返回一个页面
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	r.POST("/uploadfile", func(c *gin.Context) {
		// 从请求中读取文件, <input type="file" name="uf1" id="">
		file, err := c.FormFile("uf1")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": "error",
				"error":  err.Error(),
			})
		} else {
			// 将读取到的文件保存到本地服务器中
			log.Println(file.Filename)
			dst := fmt.Sprintf("./file/%s", file.Filename)
			c.SaveUploadedFile(file, dst)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.Run(":9090")
}
