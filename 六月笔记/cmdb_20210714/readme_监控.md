1、如何监听操作系统信号
func main() {

	// 交叉编译, set GOOS=linux go build GOOS=windows (在cmd里面设置)
	// 管道
	interrupt := make(chan os.Signal, 1)
	reload := make(chan os.Signal, 1)

	// 监听系统信号,一般只监听 SIGINT(中断信息相当于Ctrl + c) SIGHUP 这2种信号
	signal.Notify(interrupt, syscall.SIGINT)
	signal.Notify(reload, syscall.SIGHUP)

	// 等待信号,读取管道
INTERRUPT:
	for {
		select {
		case <-interrupt:
			break INTERRUPT
		case <-reload:
			fmt.Println("raload")
		}
	}
}