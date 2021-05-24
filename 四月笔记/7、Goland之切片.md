## Goland之切片

### 一、什么是切片

**切片是长度可变的数组**(具有相同数据类型的数据项组成的一组长度可变的序列)，切片由三部分组成:

- 指针：指向切片第一个元素指向的数组元素的地址

- 长度：切片元素的数量

- 容量：切片开始到结束位置元素的数量

- 切片的本质：左包含，右不包含(大于等于左边的，小于右边的)

- 总结一下就是：要从切片a中删除索引为`index`的元素，操作方法是`a = append(a[:index], a[index+1:]...)`

- ```
  a[2:]  // 等同于 a[2:len(a)]
  a[:3]  // 等同于 a[0:3]
  a[:]   // 等同于 a[0:len(a)]
  ```

```go
package main

import "fmt"

func main() {

	// 声明切片，零值是nil
	var arry []int
	fmt.Printf("%T \n", arry)

	// 字面量（赋值），这时候赋值可以是任意的
	arry = []int{1, 2, 3, 4, 5}
	fmt.Println(arry)

	arry = []int{1, 2, 3, 1, 6, 7, 8, 9}
	fmt.Println(arry)

	// 数组切片赋值
	var nums [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	test := nums[1:5]
	fmt.Printf("%#v\n", test)

	// 使用make函数，初始化切片
	var arry1 []int
	arry1 = make([]int, 3)
	fmt.Printf("%#v %d %d \n", arry1, len(arry1), cap(arry1))
	arry1 = make([]int, 4, 5)
	fmt.Printf("%#v %d %d \n", arry1, len(arry1), cap(arry1))

	// 元素操作 (CRUD)
	var arry2 []int
	arry2 = make([]int, 5, 7)
	fmt.Printf("%#v %d %d \n", arry1, len(arry1), cap(arry1))
	fmt.Println(arry2[1])
	fmt.Println(arry2[2])
	// fmt.Println(arry2[6])   // 会报错，因为第6个是nil，不可操作
	arry2[2] = 11
	fmt.Println(arry2[2])

	arry2 = append(arry2, 4)
	fmt.Printf("%#v %d %d \n", arry2, len(arry2), cap(arry2))
	fmt.Println("-------------------------")
	for i := 0; i < len(arry2); i++ {
		fmt.Println(arry2[i])
	}

	// 切片操作
	fmt.Printf("%T %#v", arry2[1:5])

	// copy   用copy去实现删除
	arry3 := []int{1, 2, 3}
	arry4 := []int{10, 20, 30, 40}
	copy(arry4, arry3)
	fmt.Println(arry4)

	arry5 := []int{10, 20, 30, 40}
	copy(arry3, arry5)
	fmt.Println(arry3)

	// 删第一个、最后一个
	arry6 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	//fmt.Println(arry6[1:])               // 删第一个
	//fmt.Println(arry6[0 : len(arry6)-1]) // 删最后一个

	// 删了3(第二个元素),利用底层切片共享原理
	copy(arry6[2:], arry6[3:])
	fmt.Println(arry6[:len(arry6)-1])

	// 应用
	// 堆栈  先进后出
	// 队列  先进先出
	queue := []int{}  // 队列
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	fmt.Println(queue)

	fmt.Println(queue[1:])

	// 堆栈,先进后出
	stack := []int{}
	stack = append(stack,1)
	stack = append(stack,2)
	stack = append(stack, 3)
	fmt.Println(stack,"堆栈")
	fmt.Println(stack[:len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[:len(stack)-1])
	stack = stack[:len(stack)-1]
	fmt.Println(stack[:len(stack)-1])

	// 多维切片
	points1 := [][] int{}
	points2 := make([][]int , 0)
	points1 = append(points1, []int{1,2,3})
	points1 = append(points1, []int{4,5,6})
	points2 = append(points2,[]int{7, 8, 9,})
	fmt.Println(points1)
	fmt.Println(points2)
}
```

