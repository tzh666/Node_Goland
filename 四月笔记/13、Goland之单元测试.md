## Goland之单元测试

### 一、单元测试

```go
// 1、测试覆盖率
D:\Golang_student\Golang_code\day05_pck\src\AAA> go test AAA/B -coverprofile=cover.out
ok      AAA/B   1.171s  coverage: 66.7% of statements

// 2、查看测试覆盖率
go tool cover -html cover.out
```



### 二、性能测试（基准测试）

```go
D:\Golang_student\Golang_code\day05_pck\src\AAA>go test AAA/B -bench .


D:\Golang_student\Golang_code\day05_pck\src\AAA>go test AAA/B -bench . -benchmem
```



### 三、代码结构

```sh
D:.
├─bin
├─pkg
│  └─windows_386
└─src
    ├─AAA
    │  └─B
```

![image-20210406004839249](D:\Golang_student\Golang笔记\四月笔记\13、Goland之单元测试.assets\image-20210406004839249.png)

```go
package main

import (
	"AAA/B"
	"fmt"
)

func main() {
	fmt.Println("ADD:", B.Add(1, 31))
	fmt.Println("Fact:", B.Fact(10))
}

------
package B

// 注释1
// func Add(a, c int) int {
// 	return a + c
// }

func Add(a, c int) int {
	if a > 0 {
		return a + c
	} else {
		return a
	}
}

// 递归，测试性能
func Fact(n int) int {
	if n == 0 {
		return 0
	} else {
		return n * Fact(n-1)
	}
}
-------
package B

import "testing"

//  测试对应B.go 注释1的代码
// func TestADD(t *testing.T) {
// 	if 4 != Add(1, 3) {
// 		t.Error("1 + 3 != 4")
// 	}
// }

// 测试覆盖度
func TestADD(t *testing.T) {
	if 4 != Add(1, 3) {
		t.Error("1 + 3 != 4")
	}
}

func TestAddFlag(t *testing.T) {
	// 这里的条件根据函数的返回值去写，每一个分支对应一个测试函数，就是传入的参数跟
	if -1 != Add(-1, 3) {
		t.Error("-1 + 3 != -1")
	}
}

// 基准测试
func BeanchmarkFact(b *testing.B) {
	for i := 0; i < 1100; i++ {
		Fact(i)
	}
}
```

