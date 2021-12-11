package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 解析模板,根模板得写到前面
	tpl := template.Must(template.ParseFiles("views/base.html", "views/index.html"))
	// 渲染模板
	msg := "index"
	tpl.ExecuteTemplate(w, "index.html", msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	// 解析模板,根模板得写到前面
	tpl := template.Must(template.ParseFiles("views/base.html", "views/home.html"))
	// 渲染模板
	msg := "home"
	tpl.ExecuteTemplate(w, "home.html", msg)
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
