package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2、解析模板
	tpl := template.Must(template.ParseFiles("tmp/hello.html"))

	// 3、渲染模板(其他详情参考git笔记)
	name := "欢少"
	tpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello)

	// 1、监听端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
