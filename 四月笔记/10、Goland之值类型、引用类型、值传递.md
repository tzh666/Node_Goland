## Goland之值类型、引用类型、值传递

### 一、值类型、引用类型

**值类型：  int  bool float array 指针  结构体**
**引用类型：slice map  接口  （赋值之后，改变赋值的对象，会影响到原来的值）**

```go
package main

import "fmt"

func main() {
	array := [3]string{"A", "B", "C"}
	slice := []string{"a", "b", "c"}

	arrayA := array
	sliceB := slice

	array[0] = "Z"
	sliceB[0] = "z"

	// 改变切片的值，会影响到原来的切片，因为他公用一块内存
	fmt.Println(arrayA, array)
	fmt.Println(sliceB, slice)
}
```



###  二、值传递

![image-20210404010435877](D:\Golang_student\Golang笔记\四月笔记\10、Goland之值类型、引用类型、值传递.assets\image-20210404010435877.png)

```go
package main

import "fmt"

func changInt(a int) {
	a = 100
}

func changSlice(s []int) {
	s[0] = 100
}

func changIntP(p *int) {
	*p = 100
}

func main() {
	num := 1
	changInt(num)
	fmt.Println(num) // num的值不变

	stm := []int{1, 2, 3}
	changSlice(stm)
	fmt.Println(stm) // stm的值会变

	nump := 1
	changIntP(&nump)
	//  取地址用&，获取变量的值用*
	fmt.Println(nump) // nump的值会变
}
```

