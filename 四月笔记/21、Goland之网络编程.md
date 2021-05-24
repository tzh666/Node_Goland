## Goland之网络编程

### 一、https网络编程

#### 1.1、server端

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func test02(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %d", time.Now().Unix())
}

// 处理器
type Test03 struct{}

func (t Test03) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now().Format("2006-01-02 15:04:05"))
}

func main() {

	// 定义处理器函数（面向过程）
	http.HandleFunc("/test01/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi kk"))
	})

	// 自定义接口（面向过程）
	http.HandleFunc("/test02", test02)

	// 第二个参数可以传指针类型的，也可以不穿指针类型的
	http.Handle("/test03", Test03{})  // 面向对象
	http.Handle("/test04", &Test03{}) // 面向对象

	http.HandleFunc("/request/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL, r.Proto, r.Method)
		fmt.Println(r.Referer())
		fmt.Println(r.Header)

		fmt.Println("请求体：")
		// bytes := make([]byte, 1024)
		// n, _ := r.Body.Read(bytes)
		// fmt.Println(n)
		io.Copy(os.Stdout, r.Body) // 等价于上面的三行代码

		w.Write([]byte("Request"))
	})

	// 启动文件服务器，传递个路径进去(可以用于文件的copy，如果不能ssh的话)
	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe("0.0.0.0:8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
```

### 1.2、client01端

```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	urll "net/url"
	"os"
)

func main() {
	url := "http://127.0.0.1:8888/request/"

	// get获取
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	} else {
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	}

	// post 获取
	josn := bytes.NewReader([]byte(`{"name":"kk","password":"123456"`))
	response, err = http.Post(url, "/app/json", josn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Proto)
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	}

	// PostForm 获取
	params := make(urll.Values) // 这是一个映射map
	params.Add("name", "kk")
	params.Add("password", "123456")

	response, err = http.PostForm(url, params)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Proto)
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	}
}
```

### 1.3、client02端

```go
package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	addr := "https://www.baidu.com/?tn=94819464_hao_pg"

	// DELETE 请求
	request, _ := http.NewRequest("DELETE", addr, nil)

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},// 忽略自签名https证书的校验
	}

	client := &http.Client{Transport: transport}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Proto)
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	}
}
```

### 1.4、client03端

```go
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	addr := "https://www.baidu.com/?tn=94819464_hao_pg"

	// DELETE 请求
	request, _ := http.NewRequest("DELETE", addr, nil)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response.Proto)
		io.Copy(os.Stdout, response.Body)
		fmt.Println()
	}
}
```



### 二、TCP网络编程

#### 2.1、服务端

```go
func main() {
	// 定义传输协议，端口
	addr := ("0.0.0.0:9999")
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer listener.Close()
	fmt.Println("Listen: ", addr)

	// 客户端发生信息给服务器
	client, err := listener.Accept()
	if err == nil {
		client.Write([]byte(time.Now().Format("2006-01-02 15:04:05")))
		// 关闭客户端、服务器
		client.Close()
	}
}
```

