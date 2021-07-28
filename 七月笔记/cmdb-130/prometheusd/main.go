package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"prometheusd/client"
	"prometheusd/config"
	"prometheusd/tasks"

	"github.com/imroc/req"
	"github.com/sirupsen/logrus"
)

/*
	1.配置
	2.随机uuid
	  a."github.com/google/uuid"
	  b.检查文件,如果有的话就不更新,没有的话就保存
	3.http请求 req包(https://gowalker.org/github.com/imroc/req)
	4.更新配置yaml文件
	5.应用配置
	  a.检测配置文件 promtool check
	  b.覆盖文件
	  c.热加载
*/

func main() {
	// 1.配置
	var (
		server = ""
		help   bool
		h      bool
	)
	flag.StringVar(&server, "server", "http://localhost:8888", "cmdb server")
	flag.BoolVar(&help, "help", false, "help")
	flag.BoolVar(&h, "h", false, "h")
	// 更改默认提示
	flag.Usage = func() {
		// 自定义默认提示
		fmt.Println("usage: Prometheusd --server http://localhost:8888")
		flag.PrintDefaults()
	}

	// 上面定义了几个简单的参数，在所有参数定义生效前，需要使用flag.Parse()来解析参数
	flag.Parse()

	if h || help {
		flag.Usage()
		// 退出
		os.Exit(0)
	}
	// 请求的时候开启debug模式
	req.Debug = true
	// 2.随机uuid
	option, err := config.NewOption(server)
	if err != nil {
		logrus.Fatal(err)
	}
	client := client.NewClient(option)
	configTask := tasks.NewConfigTask(option, client)
	heartbeatTask := tasks.NewHeartbeatTask(option, client)
	registertTask := tasks.NewRegistertTask(option, client)

	// 3.通信  (心跳,注册)
	go configTask.Run()
	go heartbeatTask.Run()
	go registertTask.Run()

	// 信号量,控制程序运行不退出
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGABRT)
	// 写
	<-interrupt
}
