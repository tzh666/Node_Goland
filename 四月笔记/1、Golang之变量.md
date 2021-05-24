## Golang之变量

### 一、变量定义规范

```shell
# 变量名需要满足标识符命名规则
1.必须由非空的unicode字符串组成、数字、_
2.不能以数字开头
3.不能为go的关键字(25个)
4.避免和go预定义标识符冲突，true/false/nil/bool/
5.驼峰命名法
6.标识符区分大小写
```



### 二、变量声明方式

```shell
# 如果声明的时候不赋值，然后打印，那么string类型变量默认给0
# 声明方式1
var me string
me = "666"

# 声明方式2
var me string = "666"

# 声明方式3
var name,user string = "张三","admin"

# 声明方式4
var (
		age    int = 1
		height string = "A"
	)

# 声明方式5，推荐使用
var (
		age     = "A"
		height  = "A"
	)

# 声明方式6，简短声明。只能在函数内部使用
isBoy := false
```



### 三、变量类型

####  3.1、布尔类型

布尔类型用于表示真假，类型名为bool，只有两个值true和false，占节宽度，**零值为false**

```shell
package main

import "fmt"

func main() {

	//  定义
	var zero bool
	isBoy := true
	isGirl := false
	fmt.Println(zero, isBoy, isGirl)

	// 操作
	fmt.Println(true && false)
	
	fmt.Println(true && true)
	fmt.Println(false && false)
	
	fmt.Println(false || false)
	
	fmt.Println(!isBoy)
}
```

#### 3.2、整型

Go语言提供了`5种有符号、5种无符号、1种指针、1种单字节、1种单个unicode字符(unicode码点）`，共13种整数类型,**零值均为0**.
`int, uint, rune, int8, int16, int32, int64, uint8, uint16, uint32, uint64,byte, uintptr`

```shell
package main

import "fmt"

func main() {
	var age int = 24
	// %T 打印变量的类型，%d 占位符 \n 换行
	fmt.Printf("%T %d\n", age, age)

	// 操作
	fmt.Println(1 + 2)
	fmt.Println(2 - 1)
	fmt.Println(2 * 1)
	fmt.Println(2 / 1)

	// ++ -- 只能在后面，不能在前面
	var test1 = 1
	test1++
	fmt.Println(test1) // 2

	var test2 = 2
	test2--
	fmt.Println(test2) // 1

	// 关系运算
	fmt.Println(2 > 3)
	fmt.Println(2 == 3)
	fmt.Println(2 != 3)

	// 位运算  二进制的运算 10 => 2

	// 赋值运算，=，+=，-=，*=，/=，%=，&=，|=，^=
	var aa = 1
	aa += 3
	fmt.Println(aa)
}
```

#### 3.3、浮点数

浮点数有float32、float64的，默认是64的（float），常用的也是64的。**零值的话也是0！**

```go
package main

import "fmt"

func main() {
	var age float64 = 24.1

	// 查看%F查看数据类型  %f float64的占位符
	fmt.Printf("%F %f", age, age)

	// 操作 跟整数的一样
}
```

#### 3.4、字符串

```go
package main

import "fmt"

func main() {

	// 定义方式一，可解释字符串, \转义
	var name1 string = "k\tk"
	fmt.Println(name1)

	// 定义方式二，原生字符串，写正则用的多这种方式
	// 特殊字符 \t \r \n \f \b
	var name2 string = `奥里\t给`
	fmt.Println(name2)

	var name3 string = "k\\tk"
	fmt.Println(name3)

	// 操作
	// 字符串连接：+
	var a = "xxx"
	var b = "yyy"
	fmt.Println(a + b)

	// 关系运算（== ！= > >= < <=  ,只比较第一个字母
	fmt.Println("bb" > "aaab")

	// 赋值运算
	s := "我叫"
	s += "kk"
	fmt.Println(s)

	// 索引,索引跟python的类似 从0-1
	desc := "abcdef"
	fmt.Printf("%T %C\n", desc[0], desc)

	// 切片，只能对ASCII码的字符串进行切片，不支持中文
	// [start:end-1]
	fmt.Println(len(desc))
	fmt.Printf("%T %s\n", desc[1:2], desc[1:2])
}
```

