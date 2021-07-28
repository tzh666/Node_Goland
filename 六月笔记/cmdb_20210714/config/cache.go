package config

import (
	"fmt"
	"github.com/astaxie/beego/cache"
)

var Cache cache.Cache

func Init(adapter, config string) {
	var err error
	// 这里不能用: 用了就重新定义了一个Cache，就继承不到cache.Cache
	Cache, err = cache.NewCache(adapter, config)
	if err != nil {
		fmt.Println(err)
	}
}
