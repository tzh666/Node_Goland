## Goland之接口

### 一、接口的简单使用

```go
package main

import "fmt"

// 接口是自定义类型，是对其他类型行为的抽象
// 定义一个接口
type Sender interface {
	Send(to, msg string) error
	SendAll(tos []string, msg string) error
}

// 定义结构体
type EmailSender struct {
}

// 定义结构体函数Send，跟接口的Send函数相同（函数名、参数类型、返回值）
func (s EmailSender) Send(to, msg string) error {
	fmt.Println("发送信息给：", to, "消息内容是：", msg)
	return nil
}

// 定义结构体函数SendAll，跟接口的Send函数相同（函数名、参数类型、返回值）
func (s EmailSender) SendAll(tos []string, msg string) error {
	for _, to := range tos {
		// 方法调用方法
		s.Send(to, msg)
	}
	return nil
}

func main() {

	// 把结构体赋值给接口，两个函数（函数名、参数类型、返回值）相同，所以可赋值
	var sender Sender = EmailSender{}

	// 调用接口
	sender.Send("kk", "晚上好")
	sender.SendAll([]string{"kk", "ll", "qq"}, "凌晨")
}
```



### 二、接口嵌套（匿名接口）

```go
package main

import "fmt"

type Sender interface {
	Send(msg string) error
}

type Reciver interface {
	Recive() (string, error)
}

type Client interface {
	// 接口的匿名嵌入
	Sender
	Reciver
	// 当然，接口的匿名嵌入，接口嵌入了其他接口的同时，也可以有自己的接口，结构体赋值给接口的时候，则需要实现对应的方法
	Close() error
}

// 新建一个结构体，然后新建2个方法，方法名要跟接口的一样
type MSNClient struct{}

func (c MSNClient) Close() error {
	fmt.Println(" I'm Close ")
	return nil
}

func (c MSNClient) Send(msg string) error {
	fmt.Println("Send", msg)
	return nil
}

func (c MSNClient) Recive() (string, error) {
	fmt.Println("Recive")
	return " ", nil
}

func main() {

	msn := MSNClient{}
	// 使用匿名接口
	var s Sender = msn
	var r Reciver = msn
	var c Client = msn

	s.Send("111")
	r.Recive()
	c.Send("222")
	c.Recive()
	c.Close()

	// 接口变量
	var clonser interface {
		Close() error
	}
	clonser = msn
	clonser.Close()
}
```



### 三、空接口

```go
package main

import (
	"fmt"
)

type EStruct struct{}

type Empty interface{}

// 空接口类型，可以接受任意参数
func fargs(args ...interface{}) {
	fmt.Println("------------------------------------")
	for _, arg := range args {
		switch v := arg.(type) {
		case int:
			fmt.Printf("Int: %T %v\n", v, v)
		case string:
			fmt.Printf("string: %T %v\n", v, v)
		default:
			fmt.Printf("其他的: %T, %v \n", v, v)
		}
	}
}

func main() {
	es := EStruct{}
	var e Empty

	fmt.Println(es, e)
	fargs("xxx", 123, 11.11)
}
```



### 四、反射

```go

```

