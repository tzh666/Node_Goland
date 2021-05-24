# Goland之函数

### 一、函数的多种定义方式

```go
package main

import "fmt"

// 定义一：无参无返回值的函数
func test() {
	fmt.Println("你好")
}

// 定义二：有参无返回值函数（形参）
func sayHi(name string) {
	fmt.Println("我的名字是:", name)
}

// 定义三：有参数有返回值 int定义的是返回值的类型
func sum(a, b int) int {
	return a + b
}

func main() {

	// 调用函数（调用）
	test()

	// 调用函数（实参）
	sayHi("tom")

	// 调用
	tmp := sum(1, 2)
	fmt.Println(tmp)
}
```



### 二、传参的多种方式

```go
package main

import "fmt"

// 方式一，类型一样的参数，保留最后一个的类型即可
func add(a, b, c int, d, f string) int {
	fmt.Println("我是：", d, f)
	return a + b + c
}

// 方式二，... 传可变参数（n个），只能有一个可变参数！类型必须是一致的！
func addN(a, b int, args ...int) int {

	// fmt.Println(a, b, args)
	fmt.Printf("%T \n", args) // 是一个切片
	tmp := a + b
	for _, v := range args {
		tmp += v
	}
	return tmp
}

// 方式三，可变参数，可变参数的解包
func calc(op string, a, b int, args ...int) int {
	switch op {
	case "add":
		return addN(a, b, args...) // 解包
	}
	return -1
}

func main() {
	fmt.Println(addN(1, 2, 7, 8))

	//  调用解包，函数
	args := []int{2, 2, 3, 4, 5}
	fmt.Println(addN(1, 2, args...))
	fmt.Println(calc("add", 1, 2, args...))
}
```



### 三、拆包的应用

```go
// 解包的应用，删除切片中的某一个元素
	delete_splic := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 例如删除下标为2的元素，原理是先获取0-1的元素，然后或者3以后的元素，讲这个2个切片相加即可
	fmt.Println(delete_splic[:2]) // 获取 元素1、2
	fmt.Println(delete_splic[3:]) // 获取下标为3以后的元素
	new_splic := append(delete_splic[:2], delete_splic[3:]...)  //获取下标为3以后的元素 然后通过解包的形式传递进去（append 可以同时添加多个元素）
	fmt.Println(new_splic)
```



### 四、函数的返回值

`有多个return只会返回第一个return的值`，**但是支持多个返回值**

```go
package main

import "fmt"

// 建议使用
func addNn(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b

	// 因为上面有个返回值，所以这个return不会执行到这里
	return -1, -2, -3, -4
}

// 命名返回,少用
func addNn1(a, b int) (c, d, f int) {
	return
}

func main() {
	// 返回多少个值就用多少个值接收，不想要就 _ 接收
	a, b, c, d := addNn(2, 2)
	fmt.Println(a, b, c, d)

	fmt.Println(addNn1(2, 3))
}
```



### 五、递归函数

`函数自己调用函数自己`

`得有终止的条件`

```go
package main

import "fmt"

func dg(n int) int {

	if n <= 0 {
		return -1
	} else if n == 1 {
		return 1
	}

	return n + dg(n-1)
}

func main() {
	fmt.Println(dg(5))
}
```



### 六、函数类型（匿名函数）

```go
package main

import "fmt"

func add1(a, b int) int {
	return a + b
}

func printt(callback func(...string), args ...string) {
	fmt.Printf("printt 函数输出")
	callback(args...)
}

// 然后把list函数，当作参数传递给printt函数的第一个参数，因为类型相同所以可以传递
func list(agrs ...string) {
	for k, v := range agrs {
		fmt.Println(k, ":  ", v)
	}
}

func main() {

	fmt.Printf("%T", add1) // 此时的add类型是 func(int, int) int
	// 所以就可以进行赋值,类型一致才可以赋值
	var f func(int, int) int = add1
	fmt.Printf("%T", f)

	// 调用
	printt(list, "A", "B", "C")

	// 匿名函数，得写在main函数里面
	// 定义方式一
	sayHello := func(name string) {
		fmt.Println("你好", name)
	}

	// 定义方式二，定义且直接调用
	func(name string) {
		fmt.Println("你好", name)
	}("kk")

	// 匿名函数调用
	sayHello("振欢")
}
```



### 七、闭包

闭包说的是变量的声明周期

```go
package main

import (
	"fmt"
)

func main() {

	// 定义在main函数里面
	addBase := func(base int) func(int) int {
		// 返回值为函数
		return func(n int) int {
			return base + n
		}
	}
	// 调用
	test := addBase(5) // 此时返回值是一个函数
	fmt.Printf("%T \n", test)
	//  所以还可以调用一次
	test1 := test(5) // 此时返回值是整数
	fmt.Printf("%T \n", test1)
	println(test1) // test1 = 10

	// 最终的调用方式为
	A := addBase(5)(5)
	fmt.Println(A)
}
```

