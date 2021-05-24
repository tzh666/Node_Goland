## Goland之序列化和反序列化

### 一、序列化

`内存--->字符串/字节切片`

```go
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	names := []string{"珊珊", "夸克", "欢欢"}
	users := []map[string]string{{"name": "111", "addr": "aaa"}, {"name": "222", "addr": "bbb"}}
	// 序列化
	bytes, err := json.Marshal(names)
	if err == nil {
		// fmt.Println(bytes, len(bytes))
		fmt.Println(string(bytes))
	}
    
	bytes, err = json.Marshal(users)
	if err == nil {
		// fmt.Println(bytes, len(bytes))
		fmt.Println(string(bytes))
	}
    
	// MarshalIndent   json格式转换显示
	bytes, err = json.MarshalIndent(users, "", "\t")
	if err == nil {
		// fmt.Println(bytes, len(bytes))
		fmt.Println(string(bytes))
	}
}
```



### 二、反序列化

`字符串/字节切片--->内存`

```go
names02 := []string{}
	// json.Unmarshal返回值是一个错误信息，如果成功那么就会把内存的中数据（bytes），序列化到切片names02中
	// 不能转换json的
	err = json.Unmarshal(bytes, &names02)
	fmt.Println(err)
	fmt.Println(names02)
```



### 三、json判断格式

```go
// 判断json格式是否准确
	fmt.Println(json.Valid([]byte("[]"))) // ture
	fmt.Println(json.Valid([]byte("[]x"))) // false
```



### 四、反序列化结构体

```go

```

