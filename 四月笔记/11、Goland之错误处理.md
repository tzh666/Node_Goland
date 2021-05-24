## Goland之错误处理

### 一、错误处理

```go
package main

import (
	"errors"
	"fmt"
)

// 在go里面，推荐用户自定义错误类型?    error
// 怎么创建错误类型的值?
// 无错误 返回nil

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("division by zero")
	}
	return a / b, nil
}

func main() {
	fmt.Println(division(1, 3))

	// 方式一
	if v, err := division(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	// 方式二,使用fmt.Errorf
	e := fmt.Errorf("Error: %s", "division by zero")
	fmt.Printf("%T, %v\n", e, e)
}
```



### 二、延迟执行

```go
package main

import "fmt"

func main() {

	// defer main函数退出之前执行，也就是最后执行吧
	defer func() {
		fmt.Println("defer")
	}()

	// 后进先出，所以这个defer先执行
	defer func() {
		fmt.Println("defer111")
	}()

	// 会先比defer先执行
	fmt.Println("我会先比defer先执行")

}
```

