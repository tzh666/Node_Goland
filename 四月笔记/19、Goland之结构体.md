## Goland之结构体

### 一、结构体的定义

```go
package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Reake    string
}

func main() {
	var hi User
	fmt.Printf("%T\n", hi)
	fmt.Printf("%#v \n", hi)

	me := User{1, "kk", time.Now(), "xx", "xxx", "xxxx"}
	fmt.Println(me)

	me2 := User{
		2,
		"22",
		time.Now(),
		"xxx",
		"1",
		"1",
	}
	fmt.Printf("%#v \n", me2)

	// 结构体的指针类型
	var pointer *User
	fmt.Printf("%T\n", pointer)
	fmt.Printf("%#v \n", pointer)
}
```



### 二、结构体的简单使用

```go
package main

import (
	"fmt"
	"time"
)

type User struct {
	ID       int
	Name     string
	Birthday time.Time
	Addr     string
	Tel      string
	Reake    string
}

func main() {
	me := User{
		1,
		"tzh",
		time.Now(),
		"B",
		"C",
		"A",
	}
	// 访问结构体某个属性
	fmt.Println(me.Name)

	// 修改结构体某个属性的值
	me.Name = "tzh666"
	fmt.Println(me.Name)

	// 指针结构体
	me2 := &User{
		ID:   2,
		Name: "qaz",
	}
	// 访问指针属性
	fmt.Println(me2.ID)
	// 修改指针属性（语法糖）
	me2.Birthday = time.Now()
	fmt.Println(me2.Birthday)
}
```



### 三、匿名结构体

```go
package main

import "fmt"

func main() {

	// 匿名结构体
	var me struct {
		ID   int
		Name string
	}
	fmt.Printf("%T \n", me)
	fmt.Printf("%#v \n", me)
	me.Name = "kk"
	fmt.Println(me.Name)

	// 匿名结构体,只声明一次
	// 应用场景：项目的配置，web开发的时候
	me2 := struct {
		ID   int
		Name string
	}{1, "kk"}

	fmt.Println(me2)
}
```



### 四、组合结构体

```go
package main

import "fmt"

type Address struct {
	R  string
	ss string
	No string
}

type User struct {
	ID   int
	Name string
	Addr Address   // 嵌入
}

//  结构体嵌入
func main() {
	me03 := User{
		ID:   2,
		Name: "kk",
		Addr: Address{
			R:  "北京",
			ss: "xx1",
			No: "xx2",
		},
	}

	fmt.Println(me03)
}

```



### 五、匿名组合嵌入

```go
package main

import "fmt"

type Address struct {
	R  string
	ss string
	No string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

// 匿名结构体组合
type Emp struct {
	User
	Salary float64
}

func main() {
	var me Emp
	fmt.Printf("%T, %#v \n", me, me)

	me02 := Emp{
		User: User{
			ID:   1,
			Name: "kk",
			Addr: Address{"xx", "xx", "xx"},
		},
		Salary: 111111,
	}
	fmt.Println(me02)

	// 访问属性
	fmt.Println(me02.Addr.ss)
    fmt.Println(me02.User.ID) // fmt.Println(me02.ID)  可以直接使用这种(会优先寻找Emp是再去找User的，如果直接想找具体的就写me02.User.ID)
}
```



###  六、组合指针结构体

```go
package main

import "fmt"

type Address struct {
	R  string
	ss string
	No string
}

type User struct {
	ID   int
	Name string
	Addr *Address   //指针
}

//  结构体嵌入
func main() {
	me03 := User{
		ID:   2,
		Name: "kk",
		Addr: &Address{   // 获取指针的值
			R:  "北京",
			ss: "xx1",
			No: "xx2",
		},
	}

	fmt.Println(me03.Addr.R)
}
```



### 七、匿名指针组合结构体

```go
package main

import "fmt"

type Address struct {
	R  string
	ss string
	No string
}

type User struct {
	ID   int
	Name string
	Addr Address
}

// 匿名结构体组合
type Emp struct {
	*User  // 指针
	Salary float64
}

func main() {
	var me Emp
	fmt.Printf("%T, %#v \n", me, me)
	me2 := Emp{
		User: &User{
			ID:   1,
			Name: "xxx",
			Addr: Address{
				"xx",
				"xx",
				"xx",
			},
		},
		Salary: 11111,
	}
	fmt.Println(me2)
	fmt.Println(me2.User.ID)
	fmt.Println(me2.Name)

}
```



### 八、new函数（构造函数）

```go
package main

import "fmt"

type Address struct {
	R  string
	ss string
	No string
}

type User struct {
	ID   int
	Name string
	Addr *Address
}

//
func NewUser(id int, name string, R, ss, No string) User {
	return User{
		ID:   id,
		Name: name,
		Addr: &Address{R, ss, No},
	}
}

func main() {
	test := NewUser(111, "aaaa", "xxx", "xxx", "xxx")
	fmt.Println(test.Addr.No)
}
```

