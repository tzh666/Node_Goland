package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {

	/*
		主要的:
			exec.Command()  Output() cmd.ProcessState.ExitCode()
			管道就写到shell脚本,然后go调用脚本执行
	*/

	// 查看文件路径---等价于--> which ping
	fmt.Println(exec.LookPath("ping"))

	// 第一个参数是执行的命令名称,后面的都是参数(ping baidu.com -c 5)
	cmd := exec.Command("ping", "baidu.com", "-c", "5")
	stdout, _ := cmd.StdoutPipe()

	// 异步执行
	cmd.Start()
	fmt.Println("started")

	io.Copy(os.Stdout, stdout)
	cmd.Wait() // 等待进程执行结束,能让ping 实时输出

	// 获取进程的退出状态码, 0 -> 是正常退出,1 -> 是错误  --> echo $?
	fmt.Println(cmd.ProcessState.ExitCode())

	// 更改进程状态码,很少用
	os.Exit(100)
}
