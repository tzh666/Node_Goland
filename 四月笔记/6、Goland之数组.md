## Goland之数组

### 一、数组

int类型的数组默认值是       **0**

bool类型的数组默认值是   **0**

string类型的数组默认值是 **" "**

数组定义方式：  

​			**var  数组名  [长度]类型** 

​			**数组名 := [长度]类型{元素}**

#### 1.1、数组的定义方式

```go
package main

import "fmt"

func main() {
	var nums1 [10]int
	var nums2 [10]bool
	var nums3 [10]string
    nums4 := [10]int {1,2,3}
	fmt.Printf("%T %d\n", nums1, nums1)
	fmt.Printf("%T %t\n", nums2, nums2)
	fmt.Printf("%T %q\n", nums3, nums3)
}
```

#### 1.2、数组赋值

```go
package main

import "fmt"

func main() {
	var nums1 [10]int
	var nums2 [10]bool
	var nums3 [10]string
	nums4 := [10]int{1, 2, 3}
	nums5 := [...]int{1, 2, 3}
	fmt.Printf("%T %d\n", nums1, nums1)
	fmt.Printf("%T %t\n", nums2, nums2)
	fmt.Printf("%T %q\n", nums3, nums3)
	fmt.Printf("%T %#v\n", nums4, nums4)
	fmt.Printf("%T %#v\n", nums5, nums5)

	// 字面量，赋值
	nums1 = [10]int{10, 20, 30}
	fmt.Println(nums1)

	nums1 = [10]int{1: 20, 9: 20}
	fmt.Println(nums1)

	nums1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(nums1)

}
```

#### 1.3、操作

```go
package main

import "fmt"

func main() {
	var nums1 [10]int
	var nums2 [10]bool
	var nums3 [10]string
	nums4 := [10]int{1, 2, 3}
	nums5 := [...]int{1, 2, 3}
	fmt.Printf("%T %d\n", nums1, nums1)
	fmt.Printf("%T %t\n", nums2, nums2)
	fmt.Printf("%T %q\n", nums3, nums3)
	fmt.Printf("%T %#v\n", nums4, nums4)
	fmt.Printf("%T %#v\n", nums5, nums5)

	// 字面量，赋值
	nums1 = [10]int{10, 20, 30}
	fmt.Println(nums1)

	nums1 = [10]int{1: 20, 9: 20}
	fmt.Println(nums1)

	nums1 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Println(nums1)

	// 操作
	// 比较数组是否相同
	arr1 := [3]int{1, 2, 3}
	arr2 := [3]int{1, 2, 4}
	fmt.Println(arr1 == arr2)

	// 获取数组长度
	fmt.Println((arr1))

	// 索引 0...len(array)-1
	fmt.Println(arr1[0])
	for i := 0; i < len(arr1); i++ {
		fmt.Println("第", i, "个：  ", arr1[i])
	}

	// 在使用  " := " 赋值，在左侧必须得有一个新的变量！否在会报错
	for _, value := range arr2 {
		fmt.Println(value)
	}

	// 数组切片,数组切片出来的类型不是数组, %v 查看元素
	arr3 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Printf("%T\n", arr3[1:3])
	fmt.Printf("%v", arr3[1:3])

	// 多维数组，掌握即可
	var marry [2][2]int
	fmt.Println
}
```

