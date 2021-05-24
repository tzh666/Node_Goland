## Goland之文件操作2

### 一、filepath

```go
func main() {
	fmt.Println(filepath.Abs(".")) // 获取当前路径

	fmt.Println(filepath.Abs(os.Args[0])) // 获取当前文件运行的绝对路径

	path, _ := filepath.Abs(os.Args[0]) // 获取当前文件运行的绝对路径
	dirPath := filepath.Dir(path)
	fmt.Println(dirPath)

	inPath := dirPath + "/conf/ip.ini" // 路径拼接
	fmt.Println(inPath)
	fmt.Println(filepath.Join(dirPath, "conf", "ip.ini")) // 路径拼接

	fmt.Println(filepath.Glob("./[abcm]*.go")) // 文件查找，当前目录以ab开头的

	// 打印当前目录所以文件的信息
	filepath.Walk(".", func(path string, fileInfo os.FileInfo, err error) error {
		fmt.Println(path, fileInfo.Name())
		return nil
	})
}
```



### 二、sysdir

```go
func main() {
	fmt.Println(os.TempDir())
	fmt.Println(os.UserHomeDir())
	fmt.Println(os.UserCacheDir())
	fmt.Println(os.Getwd())
}
```



### 三、标准输入输出

```go
func main() {
	fmt.Println("xxx")
	os.Stdout.Write([]byte("xxx")) // 这2个输出是一样的，标准输出
	fmt.Println()

	byte := make([]byte, 10)
	n, err := os.Stdin.Read(byte) // 标准输入
	fmt.Println(n, err, byte)

	scanner := bufio.NewScanner(os.Stdin) // 拿到控制台输入的，再写入
	scanner.Scan()
	fmt.Println(scanner.Text()) // 再输出
}
```



### 四、CSV文本生成

```go
func main() {
	file, err := os.Create("user.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		write := csv.NewWriter(file)
		write.Write([]string{"name", "love"})
		write.Write([]string{"kk", "qq"})
		write.Flush()
	}
}
```



### 五、CSV文本读取

```go
func main() {
	file, err := os.Open("user.csv")
	if err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		reader := csv.NewReader(file)
		for {
			line, err := reader.Read()
			if err != nil { // 先判断文件打卡是否有错误
				if err != io.EOF { // 再判断是否读取到末尾
					fmt.Println()
				}
				break // 如果err != os.EOF 则break
			} else {
				fmt.Println(line)
			}
		}
	}
}
```



### 六、文件流buffer操作

```go
func main() {
	// 内存中操作字符串
	// 应用场景：比如网络中收到一份数据，不想让数据落盘，就在流中操作再让数据落盘
	reader := strings.NewReader("11234")
	// bytes := make([]byte, 3)
	// for {
	// 	line, err := reader.Read(bytes) // 每次读取三个字节
	// 	if err != nil {
	// 		if err != io.EOF {
	// 			fmt.Println(err)
	// 		}
	// 		break
	// 	} else {
	// 		fmt.Println(line, bytes[:line])
	// 	}
	// }
	scanner := bufio.NewScanner(reader) // 一行一行读
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
```