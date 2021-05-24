## Goland之sync

### 一、sync-map

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	var users sync.Map
	users.Store(10, "kk")
	users.Store(20, "ll")

	if value, ok := users.Load(10); ok {
		fmt.Println(value.(string))
	}

	if value, ok := users.Load(20); ok {
		fmt.Println(value.(string))
	}

	users.Delete(10)
	if value, ok := users.Load(10); ok {
		fmt.Println("kk已经阵亡！！！")
		fmt.Println(value.(string))
	}
}
```



### 二、once

```go
func main() {

	var once sync.Once
	for i := 0; i < 100; i++ {
		// 用once修饰的函数只会执行一次
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
```

