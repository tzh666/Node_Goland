## Golang之常量

### 一、常量

`常量，也就是定义以后不可以改变的数据类型`

`常量我们约定都是大写字母`

```go
// 定义省略类型的常量
const NAME1 = "tzh"

// 定义多个常量类型相同
const NAME2, NAME3 string = "tzh2", "tzh3"

// 定义多个常量，类型不同
const (
	NAME4 = "tzh4"
	AGE1  = 24
	)

// 定义多个常量，类型省略
const NAME5, AGE2 = "tzh5", 24	

// 赋值相同的数字，可以使用以下方式
const (
    C1 int = 1
	C2
	C3
	C4 string = "C4"
	C5	
	)	 	
```

应用场景

```go
// 枚举
const (
	C6 int = iota
		C7
		C8
	)
```

