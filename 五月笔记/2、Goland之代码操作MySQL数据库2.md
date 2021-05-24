## Goland之代码操作MySQL数据库2

### 一、目录截图

![image-20210516183743073](D:\GO\Goland之代码操作MySQL数据库2.assets\image-20210516183743073.png)

### 二、代码

#### 2.1、main.go

```go
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库信息
const (
	dbUser     = "root"
	dbPassWord = "123456"
	dbHost     = "192.168.241.66"
	dbPort     = 3306
	dbName     = "goland"
	dbDriver   = "mysql"
)

// 服务器监听信息
const (
	Listenaddr = ":9999"
)

// sql信息
const (
	sqlTasks       = "select id, name from testkk"
	sqlCreateTasks = "insert into testkk values(?,?)"
	sqlDeleteTask  = "delete from testkk where id = ?"
)

// 用来存储数据库对应表的结构体
type Task struct {
	Id   int
	Name string
}

// 对应提交的数据
type TaskFrom struct {
	Id   string
	Name string
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&loc=PRC&parseTime=true", dbUser, dbPassWord, dbHost, dbPort, dbName)

	// 打开数据库连接池
	db, err := sql.Open(dbDriver, dsn)
	if err != nil {
		log.Fatal("数据库连接池打开失败，错误信息为：", err)
	}

	// 测试数据库连接
	if err = db.Ping(); err != nil {
		log.Fatal("数据库连接失败，错误信息为：", err)
	}

	// http 协议，老规矩（"请求路劲",处理器(或者处理器函数),）,启动监听
	// 显示任务列表
	http.HandleFunc("/", func(reponse http.ResponseWriter, request *http.Request) {
		tasks := make([]Task, 0, 20)
		rows, err := db.Query(sqlTasks)

		// 如果查询成功
		if err == nil {
			for rows.Next() {
				var task Task
				err := rows.Scan(&task.Id, &task.Name)
				if err == nil {
					// 在切片tasks末尾添加元素
					tasks = append(tasks, task)
				} else {
					log.Fatal("添加失败错误信息为：", err)
				}
			}
		}
		// 测试是否获取到数据
		//fmt.Println(tasks)

		// 指定模板文件路径
		tpl := template.Must(template.ParseFiles("views/task.html"))
		// 执行一个模板,写到reponse中
		tpl.ExecuteTemplate(reponse, "task.html", struct {
			// 要展示的数据
			Tasks []Task
		}{tasks})
	})

	// 添加功能
	http.HandleFunc("/add/", func(reponse http.ResponseWriter, request *http.Request) {
		var (
			task   *TaskFrom
			errors = make(map[string]string)
		)

		if request.Method == http.MethodGet { // 判断是否为GET请求
			// 加载模板

		} else if request.Method == http.MethodPost { // 判断是不是POST请求
			// 获取前端提交的数据，且去除空格
			id := strings.TrimSpace(request.PostFormValue("id"))
			name := strings.TrimSpace(request.PostFormValue("name"))
			// 赋值给Task
			task = &TaskFrom{
				Id:   id,
				Name: name,
			}
			// 判断长度 中文也可以判断(数据验证)
			nameLength := utf8.RuneCountInString(task.Name)
			if nameLength == 0 {
				errors["name"] = "name错误"
			}

			// 验证完成无错误
			if len(errors) == 0 {
				// 添加数据
				re, err := db.Exec(sqlCreateTasks, task.Id, task.Name)
				if err != nil {
					log.Fatal("数据插入失败:", err, re)
				}
				// 添加成功重定向到数据展示页面(http.StatusFound---->对应状态码302)
				http.Redirect(reponse, request, "/", http.StatusFound)
			}
		}
		tmpl := template.Must(template.ParseFiles("views/add_task.html"))
		tmpl.ExecuteTemplate(reponse, "add_task.html", struct {
			// 要展示的数据
			Task   *TaskFrom
			Errors map[string]string
		}{task, errors})
	})

	// 删除功能
	http.HandleFunc("/delete/", func(reponse http.ResponseWriter, request *http.Request) {
		// 根据传递过来的ID去删除数据，首先得获取ID
		id := request.FormValue("id")
		db.Exec(sqlDeleteTask, id)

		// 删除重定向到数据展示页面(http.StatusFound---->对应状态码302)
		http.Redirect(reponse, request, "/", http.StatusFound)
	})

	// 修改功能
	http.HandleFunc("/edit/", func(reponse http.ResponseWriter, request *http.Request) {
		tpl := template.Must(template.ParseFiles("views/edit_task.html"))
		tpl.ExecuteTemplate(reponse, "add_task.html", struct {
			// 要展示的数据
		}{})
	})

	// 启动监听
	http.ListenAndServe(Listenaddr, nil)
}
```

#### 2.2、html

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>陶振欢</title>
</head>

<body>
    <a href="/add/">添加</a>
    <table>
        <thead>
            <tr>
                <th>序列ID</th>
                <th>序列Name</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
            <!-- 通过模板渲染，结构体Task中的数据（也就是数据库中的数据） -->
            {{ range .Tasks }}
                <tr>
                    <!-- 得一一对应上结构体中的字段名称，否则数据不显示 -->
                    <td>{{ .Id }}</td>
                    <td>{{ .Name }}</td>
                    <td>
                        <a href="/delete/?id={{ .Id }}">删除</a>
                        <a href="/edit/?id={{ .Id }}">修改</a>
                    </td>
                </tr>
            {{ end }}
        </tbody>
    </table>
</body>

</html><!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>陶振欢</title>
</head>

<body>
    <a href="/add/">添加</a>
    <table>
        <thead>
            <tr>
                <th>序列ID</th>
                <th>序列Name</th>
                <th>操作</th>
            </tr>
        </thead>
        <tbody>
            <!-- 通过模板渲染，结构体Task中的数据（也就是数据库中的数据） -->
            {{ range .Tasks }}
                <tr>
                    <!-- 得一一对应上结构体中的字段名称，否则数据不显示 -->
                    <td>{{ .Id }}</td>
                    <td>{{ .Name }}</td>
                    <td>
                        <a href="/delete/?id={{ .Id }}">删除</a>
                        <a href="/edit/?id={{ .Id }}">修改</a>
                    </td>
                </tr>
            {{ end }}
        </tbody>
    </table>
</body>

</html>



<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <title>数据库ADD练习</title>
</head>

<body>
    <a href="/">返回列表</a>
    <form action="/add/" method="POST">
        <label>ID:</label> <input type="text" name="id" value="" />
        <label>Name:</label> <input type="text" name="name" value="" /> {{ .Errors.name }}
        <input type="submit" value="提交" />
    </form>
</body>

</html>
```

