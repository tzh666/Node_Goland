## Goland之文件操作

### 一、文件操作之read

```go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// 1、获取文件路径
	path := "test.txt"

	// 2、打开文件，返回值文件指针和错误信息
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err) // 先判断打卡文件是否有错误，如果有错误说明打卡失败，然后打印失败信息
	} else {
		bytes := make([]byte, 10000) // 声明且初始化一个切片，用于存储将要读取文件的内容

		for {
			// 3、读取文件,返回值n是长度，err是错误
			n, err := file.Read(bytes) // 去读一个文件的内容，然后存到上面定义的切片bytes中
			if err != nil {            // 判断写入切片是否失败，失败就打印错误信息
				if err != io.EOF { // 判断是是否读到尽头
					fmt.Println(err)
				}
				break
			} else {
				fmt.Println(string(bytes[:n]))
			}
		}
		// 4、关闭文件
		file.Close()
	}
}
```



### 二、文件操作之写

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// 1、指定要写入的文件路径
	path := "tzh.txt"

	file, err := os.Create(path) // 把要写入的路径穿入给Create函数，返回值是文件指针和一个错误信息
	if err != nil {
		fmt.Println(err)
	} else {
		// 2、写文本
		file.Write([]byte("x"))  // Write函数只能传递字节数组
		file.WriteString("xxxx") // WriteString可以传递字符串
	}
	file.Close()
	// 3、关闭文件
}
```



### 三、文件操作之追加

```go
package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	path := "test.log"
	/*
		三个参数分别是，
		路径 		    path
		文件的可操作 	 os.O_APPEND|os.O_CREATE(打开且创建一个文化)
		文件的权限		 os.ModePerm
		返回值     	    返回值是文件指针和一个错误信息
	*/
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		file.WriteString(strconv.FormatInt(time.Now().Unix(), 10))
		file.WriteString("\n")
		file.WriteString("ttt")
		file.WriteString("\n")
	}
	file.Close()
}



-----------------
package main

import (
	"log"
	"os"
	"time"
)

func main() {
	path := "test1.log"
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, os.ModePerm)
	if err == nil {
		log.SetOutput(file)          // 指定日志写入到文件指针file里面
		log.SetPrefix("xxxxxx")      // 设置前缀
		log.Print(time.Now().Unix()) // 指定写入的文件信息
	}
	file.Close()
}
```



### 四、文件操作之ioutil

`从go1.16开始，ioutil包已被声明废弃，请换用io包以及os包中的函数代替。`

```go
path := "test.log"
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	} else {
		bytes, _ := ioutil.ReadAll(file) // 读取全部内容
		fmt.Println(string(bytes))
	}

---------写-----------
func main() {
	bytes, err := ioutil.ReadFile("test.log") // 直接传一个文件
	fmt.Println(string(bytes), err)

	err1 := ioutil.WriteFile("test.txt", []byte("xxxxxxxxxxxxxxxxxxx"), os.ModePerm)

	fmt.Println(err1)

}
```





### 五、带缓冲io（bufio）

```go
//Scan使扫描器前进到下一个令牌，然后该令牌将
//可通过Bytes或Text方法获得。当
//到达输入端或出现错误时，扫描停止。
//在Scan返回false之后，Err方法将返回
//在扫描过程中发生，但如果是io.EOF，则出错
//将返回零。
//如果split函数返回的空值太多，则扫描会中断
//不推进输入的令牌。这是的常见错误模式
//扫描仪。
func main() {
	path := "test.txt"
	file, err := os.Open(path)
	if err == nil { // err== nil 说明读取文件无错误，则继续执行
		defer file.Close()
		i := 0
		scanner := bufio.NewScanner(file) // 传入一个文件句柄，file指针
		for scanner.Scan() {
			fmt.Println(i, scanner.Text())
            i++
		}
	}
}


--写--
func main() {

	// 文件没有会创建,有的话 会覆盖(用open函数不行的)
	file, err := os.Create("aaa.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("123456789")
	writer.Flush()
}
```



### 六、文件操作之copyfile

```go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func copyfile(src, dest string) {
	srcfile, err := os.Open(src)
	if err != nil {
		fmt.Println(err)
	} else {
		defer srcfile.Close() // 延迟关闭
		destfile, err := os.Create(dest)
		if err != nil {
			fmt.Println(err)
		} else {
			defer destfile.Close() // 延迟关闭

			// 用缓冲区的方式
			readerfile := bufio.NewReader(srcfile)
			writerfile := bufio.NewWriter(destfile)

			bytes := make([]byte, 1024*1024)
			for {
				n, err := readerfile.Read(bytes) // 读取文件到bytes切片中
				if err != nil {
					if err == io.EOF {
						break // 如果err==EOF  说明文件读取完毕，则breake
					}
				} else {
					writerfile.Write(bytes[:n]) // 读取bytes切片中的数据，写到destfile
					writerfile.Flush()
				}
			}
		}
	}
}

func main() {

	// 定义帮助信息,flag模块
	src := flag.String("s", "", "src file")
	dest := flag.String("d", "", "dest file")
	help := flag.Bool("h", false, "help")

	// 重写Usege函数
	flag.Usage = func() {
		fmt.Println(`
Usage: copyfile -s srcfile -d destfile
Options:`)
		flag.PrintDefaults() // 默认使用自己定义的帮助信息
	}

	flag.Parse() //
	if *help || *src == "" || *dest == "" {
		flag.Usage()
	} else {
		copyfile(*src, *dest)
	}
}
```



### 七、文件操作之文件夹操作

```go
package main

import "os"

func main(){

  os.Mkdir("test001", 0644)    // 创建文件夹,文件夹存在会报错

  os.Rename("test001", "test0011") // 重命名文件夹

  os.Remove("test0011")      //删除文件夹

  os.MkdirAll("a/b/c")       // 类似于 mkdir -p

  os.RemoveAll("a/b")       // 删除该目录以及该目录的子目录

}
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("x.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("文件不存在")
		} else {
			file.Close()
		}
	}
}
```

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// 不需要数组下标所以_省略，需要元素path即可
	for _, path := range []string{"xxx", "xxx", "../src"} {
		fileInfo, err := os.Stat(path) // 传递一个文件进去，可以获取到文件的大小，名字===
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			}
		} else {
			fmt.Println(fileInfo.Name())
			fmt.Println(fileInfo.Size()) // fileInfo.Size(), fileInfo.IsDir(), fileInfo.ModTime()
			fmt.Println("——————————————————————————————————————————————————————————————")
			if fileInfo.IsDir() { // 如果是文件夹则，获取文件信息
				dirfile, err := os.Open(path)

				if err == nil {
					defer dirfile.Close()

					// 打印子文件夹下全部文件信息，包括文件名，文件夹名
					// childrens, _ := dirfile.Readdir(-1) // -1 表示所有
					// for _, children := range childrens {
					// 	fmt.Println(children.Name())
					// }
					names, _ := dirfile.Readdirnames(-1)
					for _, name := range names {
						fmt.Println(name)
					}
				}

			}
		}
	}
}
```

