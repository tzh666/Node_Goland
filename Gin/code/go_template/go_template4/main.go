package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板,根模板得写到前面
	// New 使用给定名称分配新的HTML模板
	// Delims 将原来的{{}}替换为---> {[]}
	tpl := template.Must(template.New("tpl").Delims("{[", "]}").ParseFiles("views/index.html"))
	// 渲染模板
	msg := "tzh66"
	tpl.ExecuteTemplate(w, "index.html", msg)
}
func main() {
	http.HandleFunc("/index", index)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
