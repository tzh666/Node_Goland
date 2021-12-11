package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {

	// 模板函数功能(匿名函数)
	k := func(name string) (string, error) {
		return name + "牛逼666", nil
	}

	// 自定义模板函数
	funcs := template.FuncMap{
		// 这个key(tzh)就是自定义模板函数的名称,value就是函数实现的功能
		"tzh": k,
	}

	// 解析模板到文件里(自定义模板)、Must模板检查语法
	// template.New("f.html"),创建一个名字是f.html的模板对象  (New函数和ParseFiles传的名字要对应上,ParseFiles可以有多个)
	tpl := template.Must(template.New("f.html").Funcs(funcs).ParseFiles("tmp/f.html"))

	// 渲染模板(其他详情参考git笔记)
	name := "欢少"
	tpl.Execute(w, name)
}

func main() {
	http.HandleFunc("/", sayHello)

	// 监听端口
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
