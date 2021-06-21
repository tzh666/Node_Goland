package main

import "github.com/astaxie/beego/cache"
import "fmt"
import "time"

func main() {

	/*
		存在文件中(还可以存到redis中)
		CachePath  		文件路径
		FileSuffix 		文件结尾
		EmbedExpiry		清理时间
		DirectoryLevel  文件目录级别
	*/

	cache, err := cache.NewCache("file", `{
		"CachePath"      : "cache", 
		"FileSuffix"     : ".cache",
	    "EmbedExpiry"    : "60", 
		"DirectoryLevel" : "3"}`)

	if err != nil {
		fmt.Println(err)
	}
	// 如果cache存在的话,打印
	if cache.IsExist("name") {
		fmt.Println(cache.Get("name"))
	} else {
		// 不存在就设置一个
		cache.Put("name", time.Now().Format("15:04:05"), time.Minute)
	}
	fmt.Println(cache.Get("name"))
}
