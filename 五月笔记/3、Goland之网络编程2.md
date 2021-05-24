## Goland之网络编程

### 一、高级点的TCP

```go
func main() {

	/*
		步骤：
			1、启动监听  Listen
			2、接收客户端连接 Accept
			3、处理客户端连接 (数据交换)
				3.1、告诉客户端处理数据完毕
			4、关闭客户端 defer
			5、关闭服务 defer
	*/

	// 1、启动监听,指定 协议 ip 端口
	addrx := "127.0.0.1:9999"
	listener, err := net.Listen("tcp", addrx)
	if err != nil {
		log.Fatal(err)
	}

	// 5、关闭服务 defer
	defer listener.Close()

	// 循环接收客户端请求
	for {
		// 2、接收客户端连接
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		func() { //  使用匿名函数，匿名函数只会执行一次，defer才能关闭成功。否则不会关闭
			// 4、关闭客户端 defer
			defer conn.Close()
			log.Printf("client[%s] is 连接", conn.RemoteAddr())
			// 3、处理客户端连接 (数据交换)
			reader := bufio.NewReader(conn)
			scanner := bufio.NewScanner(os.Stdin)
			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					log.Println(err)
					break
				} else {
					if string(line) == "exit" {
						break
					}
					fmt.Printf("接收到数据：%s \n", string(line))
					// 告诉客户端处理完毕.回复数据
					fmt.Print("请输入回复的消息：")
					scanner.Scan()
					fmt.Fprintf(conn, "%s \n", scanner.Text())
				}
			}
		}()
	}
}


--------client-------
func main() {

	addrx := "127.0.0.1:9999"

	// 1、确定连接的ip port, 指定连接协议
	conn, err := net.Dial("tcp", addrx)
	if err != nil {
		log.Fatal(err)
	}
	// 3、连接成功后关闭
	defer conn.Close()
	log.Println("连接成功")
	// 3.1、读取server发送过来的数据
	reader := bufio.NewReader(conn)

	// 从命令行获取，用户输入的内容
	scanner := bufio.NewScanner(os.Stdin)

	for i := 0; i < 4; i++ {

		// 给服务器发送数据
		fmt.Println("请输入要发送的消息：")
		scanner.Scan()
		// 2、写入数据
		fmt.Fprintf(conn, "%s \n", scanner.Text())

		if i == 4 {
			fmt.Fprintf(conn, "exit")
		}

		// 3.2、读取server发送过来的数据
		line, _, err := reader.ReadLine()
		if err != nil {
			log.Println(err)
			break
		}
		fmt.Printf("服务端回应: %s \n", string(line))
	}
}
```



### 二、UDP协议

```go
------server端------
func main() {

	addr := ":6789"
	packageConn, err := net.ListenPacket("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer packageConn.Close()

	ctx := make([]byte, 1024)
	for {

		//  ReadFrom会把你的数据放到字节切片里面
		n, addr, err := packageConn.ReadFrom(ctx)
		if err != nil {
			log.Fatalf("read form err : %s", err)
		}
		fmt.Println("连接的ip：port是：", addr)
		fmt.Println("客户端的数据是：", string(ctx[:n]))

		// 给client端发送数据
		packageConn.WriteTo([]byte("ertyuio"), addr)
	}
}


----------client端--------
func main() {

	addr := ":6789"
	conn, err := net.Dial("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Fprintf(conn, "Time: %s", time.Now().Format("2006-01-02"))

	// 接受server数据
	ctx := make([]byte, 1024)

	n, _ := conn.Read(ctx)

	fmt.Println("server发生过来的数据是：", string(ctx[:n]))

}
```



### 三、HTTP协议(server端)

#### 3.1、处理器函数方式

```go
func main() {
	/*
		1、定义处理器函数
		参数：http.ResponseWriter(响应), *http.Request(请求)

		2、绑定URL和处理器函数关系
		http.HandleFunc(path,处理器函数)

		3、启动WEB服务
		http.ListenAndServe(addr,nil)
	*/

	// 定义处理器函数
	timeFunc := func(response http.ResponseWriter, request *http.Request) {
		now := time.Now().Format("2006-01-02")
		fmt.Fprint(response, now) // 等价于 io.WriteString(response, now)
	}

	// 绑定URL设置
	http.HandleFunc("/time", timeFunc)
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		ctx, err := ioutil.ReadFile("index.html")
		if err != nil {
			fmt.Fprint(response, "欢迎你啊！")
		} else {
			response.Write(ctx)
		}
	})

	// 启动web服务
	http.ListenAndServe(":8989", nil)
}
```

#### 3.2、处理器方式

```go
/*
	1、定义处理器
	HadleFunc interface 自定义类型(结构体)
	ServeHTTP(response http.ResponseWriter, request *http.Request)

	2、绑定URL和处理器关系
	http.Handle("path", &结构体实例)


	3、启动WEB服务
	http.ListenAndServe(addr,nil)
*/

// 定义处理器
type TimeHandleFunc struct {
}

// 实现接口函数，函数名一定得叫ServeHTTP
func (h *TimeHandleFunc) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	now := time.Now().Format("2006-01-02")
	io.WriteString(response, now)
}

func main() {

	// 绑定URL和处理器关系
	http.Handle("/time", &TimeHandleFunc{})

	// 启动服务
	http.ListenAndServe(":9998", nil)
}
```

#### 3.3、请求方式获取

```go
func main() {

	/*
		获取请求头
		注意：如果你请求的URL未绑定，那么会按照URL中最近匹配的绑定关系去处理
		比如原有：
			/tmie、/、/pwd

			你设置的 /time/xxx/xxx  ---->那么他会去找/time
	*/

	// 处理器函数
	test := func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("请求方式是：", request.Method, "请求URL是：", request.URL, "请求协议是：", request.Proto)

		// 查看返回的信息
		fmt.Println(request.Header)
		// 获取信息
		fmt.Println(request.Header.Get("User-Agent"))
	}

	// 绑定URL和处理器函数关系
	http.HandleFunc("/", test)

	// 启动服务
	http.ListenAndServe(":6789", nil)
}
```

#### 3.4、http之GET 请求--------------------------------------

##### 3.4.1、ParseForm

```go
func main() {

	/*
		GET 请求
		获取http发起请求提交的数据
		1、解析参数
		2、获取
		3、测试访问：http://localhost:9999/?a=b&a=1&c=d
	*/
	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 1、解析参数
		request.ParseForm()
		// 2、获取
		fmt.Println(request.Form)
		fmt.Println(request.Form.Get("a"))
		fmt.Println(request.Form["a"]) // 获取很多个参数
	})

	http.ListenAndServe(":9999", nil)
}
```

##### 3.4.2、FormValue

```go
func main() {

	/*
		GET 请求
		获取http发起请求提交的数据
		1、解析参数、自动获取
		2、
		3、测试访问：http://localhost:9999/?a=b&a=1&c=d
	*/
	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 1、解析参数、、自动获取
		fmt.Println(request.FormValue("a")) // 只能获取key对应的一个值
	})

	http.ListenAndServe(":9999", nil)
}
```

#### 3.5、http之POST请求

##### 3.5.1、ParseForm方式

```go
/*
	http.post
	提交数据 请求体中
	有编码格式:
		k=v&k2=v2
		上传文件：
			multipart/form-data
		application/json {"a":"1"}
		application/x-www-form-urlencoded

	请求方式为POST
		curl -XPOST "http://localhost:9999" -d"a=1&b=2"
*/

func main() {

	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 1、解析参数
		request.ParseForm()
		// 2、获取
		fmt.Println(request.Form) // 既能获取请求体的数据，也能获取url中的数据
		fmt.Println(request.Form.Get("a"))
		fmt.Println(request.PostForm) // 只能获取请求体中的数据（就是不在url中的数据）

	})
	http.ListenAndServe(":9999", nil)
}
```

##### 3.5.2、PostFormValue方式

```go
func main() {

	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 1、解析参数、自动获取
		fmt.Println(request.PostFormValue("a")) // 只能获取请求体中的数据
	})
	http.ListenAndServe(":9999", nil)
}
```

#### 3.5.3、上传文件FormFile、ParseMultipartForm

```go
func main() {

	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		/*
			// 1、解析提交内容（上传文件一定要用post提交）
			request.ParseMultipartForm(1024 * 1024)

			// 2、获取数据
			fmt.Println(request.MultipartForm)
			// 获取文件的值
			file, _ := request.MultipartForm.File["a"][0].Open()
			// 直接打印到控制台
			io.Copy(os.Stdout, file)
		*/

		// 解析提交内容，自动获取值，curl -XPOST "http://localhost:9999" -d"a=1&b=2" -F "a=main.go"
		file, header, err := request.FormFile("a")
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(os.Stdout, file)
		fmt.Println(header.Filename, header.Size, header.Header)

	})
	http.ListenAndServe(":8888", nil)
}
```

#### 3.6、json格式

```go
/*
	json
	curl -XPOST "http://localhost:8888" -H "application/json"  -d"{\"a\":1}"
*/

func main() {

	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 打印到屏幕（流中的数据执行操作一次）
		// io.Copy(os.Stdout, request.Body)
		decoder := json.NewDecoder(request.Body)
		var info map[string]interface{}

		decoder.Decode(&info)
		fmt.Println(info)
	})
	http.ListenAndServe(":8888", nil)
}
```

#### 3.7、Cookie

```go
func parseCookie(cookie string) map[string]string {

	// 定义一个映射，用于储存cookine的 k、v
	cookieMap := make(map[string]string)

	if strings.TrimSpace(cookie) == "" {
		return cookieMap
	}

	// 返回值是一个string类型的切片
	values := strings.Split(cookie, ";")

	// 然后遍历这个切片，得到一个字符串（counter=1）,然后再次分割，再次得到一个切片 [counter 1]
	for _, value := range values {

		kv := strings.Split(value, "=")
		// 格式转换后再次赋值给cookieMap
		cookieMap[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
	}

	return cookieMap
}

func main() {

	/*
		cookie 浏览器端存储
		读取浏览器中的数据 counter
		counter +=1 设置在浏览器中(counter无，设置为0+1)
	*/

	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {

		// 读取cookie（在Header头中）
		cookie := parseCookie(request.Header.Get("Cookie"))

		counter := 0
		// 字符串转int     ,cookie["counter"] 先获取map中counter对应的key
		if v, err := strconv.Atoi(cookie["counter"]); err == nil {
			counter = v
		}

		// 设置Cookie信息
		counterCookie := &http.Cookie{
			Name:     "counter",
			Value:    strconv.Itoa(counter + 1),
			HttpOnly: true,
		}

		// 把cookie 写到响应reponse中
		http.SetCookie(reponse, counterCookie)

		// 把counter写入响应reponse中
		fmt.Fprintf(reponse, "counter: %d", counter)
	})

	http.ListenAndServe(":9999", nil)
}
```

#### 3.8、重定向

```go
func main() {
	http.HandleFunc("/home/", func(reponse http.ResponseWriter, request *http.Request) {

		// 重定向
		http.Redirect(reponse, request, "/login/", 302)

		fmt.Fprintf(reponse, "首页")
	})

	http.HandleFunc("/login/", func(reponse http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(reponse, "登录页面")
	})
	http.ListenAndServe(":8888", nil)
}
```

#### 3.9、文件服务器

```go
func main() {

	// 处理器 + URL绑定
	http.Handle("/static/", http.FileServer(http.Dir("./")))

	// StripPrefix去除前缀，如果没有StripPrefix的花，它会在www/static2目录下找，现在加了StripPrefix就通过/static2/直接找到www下的文件
	http.Handle("/static2/", http.StripPrefix("/static2", http.FileServer(http.Dir("./www"))))

	// 启动web服务器
	http.ListenAndServe(":8888", nil)
}
```



### 四、HTTP协议(client端)

#### 4.1、基础库少用

```go
func main() {

	// 返回值是响应信息、err (Get请求)
	response, err := http.Get("http://127.0.0.1:9999/?a=1&b=2")
	if err != nil {
		log.Fatal("错误是：", err)
	} else {
		// 请求方式、协议
		fmt.Println(response.Proto, response.Status)
		// 打印到控制台
		io.Copy(os.Stdout, response.Body)
	}

	// POST请求
	// POST提交参数方式1
	buffer := bytes.NewBufferString(`{"a": 1}`)
	response1, err1 := http.Post("http://127.0.0.1:9999", "application/json", buffer)
	if err != nil {
		log.Fatal("错误是：", err1)
	}
	// 请求方式、协议
	fmt.Println(response1.Proto, response1.Status)
	// 打印到控制台
	io.Copy(os.Stdout, response1.Body)

	// // POST提交参数方式2
	params := url.Values{}
	params.Add("a", "1")
	params.Add("a", "2")

	response2, err2 := http.PostForm("http://127.0.0.1:9999", params)
	fmt.Println(response2, err2)
}
```

#### 4.2、req库常用

```go

```



### 五、正则表达式

#### 5.1、Compile

```go
func main() {

	// Compile 结构体(正则表达式) ^132 开头，\\d 纯数字，{后面接8位数}，数字结尾$
	reg, err := regexp.Compile("^132\\d{8}$")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(reg) // 返回值是一个正则表达式
	}

	// 匹配 MatchString
	fmt.Println(reg.MatchString("13212345678")) // true

	reg, err = regexp.Compile("132\\d{8}")

	// 替换 Replace,132开头 八个数结尾的替换掉
	fmt.Println(reg.ReplaceAllString("替换132xxxxxx", "132????????"))
	fmt.Println(reg.ReplaceAllString("替换13207845669", "132????????"))

	// 查找 Find
	fmt.Println(reg.FindAllString("查找13207841668,13207841662,111", -1))

	// 分割(通过正则表达式做分割)
	fmt.Println(reg.Split("查找13207841668,222,111", -1)) // [查找 ,222,111]

}
```

#### 5.2、基本使用

```go
func main() {

	/*
		regexp.MatchStrings    判断跟patten相比是否符合
	*/
	patten := "^123" // 以123开头
	fmt.Println(regexp.MatchString(patten, "123xxx"))
	fmt.Println(regexp.MatchString(patten, "0123xxx"))
}
```

#### 5.3、贪婪模式

```go
func main() {

	// (至少有一个数字组成)
	reg, _ := regexp.Compile("[ab][0-9]+")
	fmt.Println(reg.FindAllString("a123-45t-yhgf-dsa", -1)) // [a123]

	// 定义非贪婪模式 (至少有一个数字组成) (?U)，最小范围匹配 匹配到就不进行了
	reg, _ = regexp.Compile("(?U)[ab][0-9]+")
	fmt.Println(reg.FindAllString("a123-45t-yhgf-dsa", -1)) // [a1]
	reg.Longest()                                           // 转换为贪婪模式
	fmt.Println(reg.FindAllString("a123-45t-yhgf-dsa", -1)) // [a123]

}
```

