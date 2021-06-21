package main

import "github.com/astaxie/beego/cache"
import "fmt"
import "time"

func main() {
	// 放在在内存里, 60s检测一次,检测到失效的key就干掉
	cache, _ := cache.NewCache("memory", `{"interval" : 60}`)

	// 获取值
	fmt.Println(cache.Get("name"))

	// 放入cache
	cache.Put("name", "kk", 10*time.Second)
	fmt.Println(cache.Get("name"))

	// 删除cache
	cache.Delete("name")
	fmt.Println(cache.Get("name"))

	// 判断cache是否存在
	fmt.Println(cache.IsExist("name"))
}
