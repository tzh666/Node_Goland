## Golang之流程控制

### 一、流程控制之if、else

```go
package main

import "fmt"

func main() {
	var yes string

	fmt.Print("有西瓜卖吗?(Y/N):")
	fmt.Scan(&yes)

	fmt.Println("老婆的想法: 买十个包子")

	if yes == "Y" || yes == "y" {
		fmt.Println("一个西瓜")
	}

	fmt.Println("老公的想法：")

	if yes == "y" || yes == "Y" {
		fmt.Println("一个包子")
	} else {
		fmt.Println("十个包子")
	}
}
```

```go
package main

import "fmt"

func main() {
	var score int
	fmt.Print("请输入分数（0-100）：")
	fmt.Scan(&score)

	if score > 100 || score < 0 {
		fmt.Println("请输入正确的的成绩")
	} else if score >= 80 && score <= 100 {
		fmt.Println("成绩很OK")
	} else if score >= 60 && score < 80 {
		fmt.Println("成绩ok")
	} else {
		fmt.Println("同学请继续努力")
	}

	fmt.Println("这同学的分数为：", score)

}
```

### 二、流程控制之switch

```shell
package main

import "fmt"

func main() {
	var score int
	fmt.Println("请输入您的分数：")
	fmt.Scan(&score)
	switch {
	case score >= 80:
		fmt.Println("成绩很ok")
	case score < 80 && score >= 60:
		fmt.Println("成绩ok")
	default:
		fmt.Println("成绩不ok")
	}
}
```



### 三、for循环

```go
package main

import "fmt"

func main() {
	// 方式一，正常写法
	test := 0
	for i := 0; i <= 100; i++ {
		test += i
	}
	fmt.Println(test)

	// 方式二,while的写法
	test1 := 0
	i := 1
	for i < 101 {
		test1 += i
		i++
	}
	fmt.Println(test1)

	// 方式三，死循环的写法
	j := 0
	for {
		fmt.Println(j)
		j++
		if j > 2 {
			break
		}
	}

	// 方式四，range
	desc := "我爱中国"
	for i, ch := range desc {
		fmt.Println(i, ch)
	}
}
```



