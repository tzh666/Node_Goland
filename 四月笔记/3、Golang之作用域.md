## Golang之作用域

### 一、作用域

`作用域：定义标识符可以使用的范围`

`在Go中使用{}来定义作用域的范围`

`作用域指变量可以使用范围。go语言使用大括号显示的标识作用域范围，大括号内包含一连串的语句，叫做语句块。`

`语句块可以嵌套，语句块内定义的变量不能在语句块外使用`

```go
package main

import "fmt"

func main() {
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		outer := 22
		{
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
	}
	// fmt.Println(inner) // 这里是使用不到inner了，受作用域的限制
	fmt.Println(outer) // outer可以
}
```

