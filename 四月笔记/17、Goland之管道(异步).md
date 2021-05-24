## Goland之管道

###  一、不带缓存的管道

```go
func main() {

	// 定义
	var channel01 chan int
	fmt.Printf("%T %#v\n", channel01, channel01)

	// 初始化（不带缓冲区的管道）
	channel01 = make(chan int)
	fmt.Printf("%T %#v\n", channel01, channel01)

	// 写，在不能写入数据时会进行堵塞
	channel01 <- 1 // 使用的时候：需要有另外的例程进行读取（相对论）就相当于没有Done()产生死锁一样
	// 读，在不能读取数据时会进行堵塞
	<-channel01 // 使用的时候：需要有另外的例程进行写入（相对论）就相当于没有Done()产生死锁一样
}

-------所以得开启例程才能使用不带缓冲区的管道--------
func main() {

	// 定义
	var channel01 chan int
	fmt.Printf("%T %#v\n", channel01, channel01)

	// 初始化（不带缓冲区的管道）
	channel01 = make(chan int)
	fmt.Printf("%T %#v\n", channel01, channel01)

	// 写
	go func() {
		channel01 <- 100
	}()

	// 读
	num := <-channel01
	fmt.Println(num) // num = 100 

}
-------------管道的应用------------
func main() {
	channel := make(chan int)
	fmt.Println("start---")
	for i := 0; i < 3; i++ {
		go func(name string) {
			for i := 'A'; i < 'F'; i++ {
				fmt.Println(name, "::", string(i))
			}
			channel <- 0 // 现在是堵塞状态，等会下面代码读取了才释放。所以达到Done()的效果
		}(strconv.Itoa(i))

	}
	for i := 0; i < 3; i++ {
		<-channel
	}
}
```



### 二、带缓冲的管道

```go
func main() {
	// 带缓冲区的管道
	channel01 := make(chan int, 3)
	fmt.Printf("%T,%#v \n", channel01, channel01)

	// 带缓冲区的写(注意不能超过缓冲区的大小，否在会死锁)
	channel01 <- 1
	channel01 <- 2
	channel01 <- 3
	// channel01 <- 4 // 注意不能超过缓冲区的大小，否在会死锁)

	// 带缓冲区的读(注意一次只能读一个，先进先读)
	fmt.Println(<-channel01)
	fmt.Println(<-channel01)
	fmt.Println(<-channel01)
}


-------------带缓冲的管道的应用------------
func main() {
	// 带缓冲区的管道
	channel01 := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("<-", i)
			channel01 <- i
		}
		// 关闭管道，使得下面读取管道信息的时候有停止的标识号，从而不会堵塞
		close(channel01)
	}()
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		for e := range channel01 {
			fmt.Println("->", e)
		}
		wg.Done()
	}()
	wg.Wait()
}
```



### 三、只读只写的管道

```go
func main() {

	// 普通缓冲管道
	channel := make(chan int, 10)
	channel <- 1
	fmt.Println(<-channel)

	// 只读管道 var readonly <-chan int
	// 只写管道 var writeonly chan<- int

	// 在这个函数中只写
	func(channel chan<- int) {
		channel <- 2
	}(channel)

	// 在这个函数中只读
	func(channel <-chan int) {
		fmt.Println(<-channel)
	}(channel)
}
```



### 四、select模型

```go
func main() {
	channel01 := make(chan int, 1)
	channel02 := make(chan int, 1)

	go func() {
		channel01 <- 1
	}()

	go func() {
		channel02 <- 1
	}()

	select {
	case <-channel01:
		fmt.Print("channel01")
	case <-channel02:
		fmt.Print("channel02")
	}
}


----------------------------------------
func main() {
	channel := make(chan int, 1)
	slice := make([]int, 10)

	// 产生0-4的随机数
	for i := 0; i < 10; i++ {
		select {
		case channel <- 0:
		case channel <- 1:
		case channel <- 2:
		case channel <- 3:
		case channel <- 4:
		}

		// 把管道的值存的切片中，因为切片定义的长度，所以这里不能用append
		slice[i] = <-channel
	}
	fmt.Println(slice)
}
```



### 五、管道超时机制

```go

```

