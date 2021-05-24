## Golang之goto

### 一、goto

```go
package main

import "fmt"

func main() {
	var yes string

	fmt.Print("有西瓜卖吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法: 买十个包子")

	if yes == "Y" || yes == "y" {
        goto END  // 在底下声明跳出的地方，一般约定为大写字母+:做为完整的跳出地方
	}
	fmt.Println("买一个西瓜")
END:  // 跳出的地方
}
```

