## Goland之并发编程

### 一、例程

```go
package main

import (
	"fmt"
	"runtime"
	"time"
)

func PrintChars(name string) {
	for ch := 'A'; ch <= 'F'; ch++ {
		fmt.Printf("%s: %c\n", name, ch)
		// Gosched 让出CPU
		runtime.Gosched()
	}
}

func main() {

	// main是主例程，go是子例程。当主例程执行完毕，子例程会被停止(不管有没有执行完毕)
	go PrintChars("函数1")
	go PrintChars("函数2")
	PrintChars("函数3")
	time.Sleep(time.Second * 3)
}
```



### 二、加锁（互斥锁）

```go
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var A int
	group := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	// 定义2个匿名函数
	inrc := func() {
		defer group.Done()
		for i := 0; i < 1000; i++ {
			lock.Lock() // 加锁
			A++
			lock.Unlock() //解锁
			// 让出cpu资源
			runtime.Gosched()
		}
	}

	decr := func() {
		defer group.Done()
		for i := 0; i < 1000; i++ {
			lock.Lock()
			A--
			lock.Unlock()
			// 让出cpu资源
			runtime.Gosched()
		}
	}

	// 同时对一个数据进行操作得加锁
	for i := 0; i < 10; i++ {
		group.Add(2)
		go inrc()
		go decr()
		group.Wait()
	}
	fmt.Println(A)
}
```



### 三、例程使用2

```go
func chars(prefix string, wg *sync.WaitGroup) {
	for ch := 'A'; ch <= 'F'; ch++ {
		fmt.Printf("%s: %c\n", prefix, ch)
		// Gosched 让出CPU
		runtime.Gosched()
	}
	wg.Done()
}

func mainChars(prefix string) {
	for ch := 'A'; ch <= 'F'; ch++ {
		fmt.Printf("%s: %c\n", prefix, ch)
		// Gosched 让出CPU
		runtime.Gosched()
	}
}

func main() {

	// 技术信号量（定义等待结构体的变量）

	/*
		ADD(n), 添加n信号
		Done(), 处理完成一个信号
		Wait(), 等待计数器归零，当所有信号处理完成才结束
	*/
	var wg sync.WaitGroup // 要么定义的时候定义成指针类型，或者函数定义的时候定义成指针类型。因为使用值传递

	// 开启一个例程
	wg.Add(1)
	go chars("gorouting", &wg)

	// 主例程
	mainChars("tzh")

	fmt.Println("等待")
	wg.Wait()
	fmt.Println("结束")
}
```

```go
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// 添加例程,没执行一次就开启一次例程
		wg.Add(1)
		// 开启例程
		go func(i int) {
			fmt.Println(i)
			// 例程执行结束就Done
			wg.Done()
		}(i)
	}
	// 等待例程结束
	fmt.Println("等待")
	wg.Wait()  // 等待例程结束,下面的代码才会执行
	fmt.Println("结束")
}
```



### 四、并发锁2

```go
// 开启一个随机数种子
func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

	wg := &sync.WaitGroup{} // 结构体，传递的时候传递指针
	lock := &sync.Mutex{}   // 结构体，传递的时候传递指针

	var a, b, count = 10000, 10000, 1000

	wg.Add(2) // 开启两个例程

	go func() {
		for i := 0; i < count; i++ {
			// 生成小于100的随机数
			money := rand.Intn(100)
			if a > money {
				lock.Lock() // 加锁
				a -= money
				// 切换例程
				time.Sleep(time.Microsecond)
				b += money
				lock.Unlock() //解锁
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < count; i++ {
			// 生成小于100的随机数
			money := rand.Intn(200)
			lock.Lock() // 加锁
			if b > money {
				b -= money
				// 切换例程
				time.Sleep(time.Microsecond)
				a += money
			}
			lock.Unlock() //解锁
		}
		// 一次完整的例程结束后就Done
		wg.Done()
	}()
	fmt.Println("等待")
	wg.Wait() // 等待例程结束,下面的代码才会执行
	fmt.Println("结束")
	fmt.Println(a, b, a+b)
}
```

```go
func main() {

	wg := &sync.WaitGroup{}
	lock := &sync.Mutex{}
	count, tmp, celi := 5, 0, 1000
	for i := 0; i < count; i++ {
		wg.Add(2)
		go func() {
			for i := 0; i < celi; i++ {
				lock.Lock()
				tmp++
				lock.Unlock()
				time.Sleep(time.Millisecond)
			}
			wg.Done()
		}()

		go func() {
			for i := 0; i < celi; i++ {
				lock.Lock()
				tmp--
				lock.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(tmp)
}
```



### 五、pool

```go
type Connection struct {
	id   int
	addr string
}

func main() {
	addr := "192.168.1.120"
	index := 0
	pool := sync.Pool{
		New: func() interface{} {
			index++
			fmt.Println("new", index)
			return &Connection{index, addr}
		},
	}
	c := pool.Get()
	fmt.Println(c)
	pool.Put(c)
	// 当Put c 回去，再次Get 不会再创建一个对象，而是用原来的 index 1
	c2 := pool.Get()
	fmt.Println(c2)
}
```

