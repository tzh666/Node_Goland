## Goland之网络编程-爬虫

### 一、标签选择器

```go
go mod init xxxx 
go run main.go
```

```go
package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {

	s := "func"
	url := "https://www.runoob.com/?s" + s

	// 发起http请求，获取响应并创建Document结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// 符合选择器
	// tag + class

	// 子孙选择器,先找符合1的再找2的再找3的
	// selector1  selector2 selector3

	selection := document.Find(".archive-list a")
	selection.Each(func(index int, tag *goquery.Selection) {

		href, exists1 := tag.Attr("href")
		fmt.Println(href, exists1)

		// 获取href 对应的value
		fmt.Println(tag.Text())
	})

}
```



### 二、类选择器

```go
go mod init xxxx 
go run main.go
```

```go
func main() {

	s := "func"
	url := "https://www.runoob.com/?s" + s

	// 发起http请求，获取响应并创建Document结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// class选择器
	// .className
	selection := document.Find(".archive-list").Find("a")
	selection.Each(func(index int, tag *goquery.Selection) {

		href, exists1 := tar.Attr("href")
		fmt.Println(href, exists1)

		// 获取href 对应的value
		fmt.Println(tar.Txet())
	})
}
```



### 三、ID选择器

```go
func main() {

	s := "func"
	url := "https://www.runoob.com/?s" + s

	// 发起http请求，获取响应并创建Document结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// id选择器
	// #id
	document.Find("#footer").Attr("class")
	document.Find("#footer").Html()
	document.Find("#footer").Text()
}
```



### 四、子孙选择器

```go
func main() {

	s := "func"
	url := "https://www.runoob.com/?s" + s

	// 发起http请求，获取响应并创建Document结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// 符合选择器
	// tag + class

	// 子孙选择器,先找符合1的再找2的再找3的
	// selector1  selector2 selector3

	selection := document.Find(".archive-list a")
	selection.Each(func(index int, tag *goquery.Selection) {

		href, exists1 := tar.Attr("href")
		fmt.Println(href, exists1)

		// 获取href 对应的value
		fmt.Println(tar.Txet())
	})

}
```

