## Goland之框架beego

### 一、MVC思想

```go
M-------> Model：模型
V-------> View：视图
C-------> Controller：控制器
```



### 二、beego之路由

#### 2.1、静态路由

```go
// 导入beego的包
import (
	"fmt"

	"github.com/astaxie/beego/context" // 导入beego的包

	"github.com/astaxie/beego" // 导入beego的包
)

func main() {

	// 路由绑定函数,  路径，处理器函数（处理器）
	// 以GET方式请求，通过绑定函数处理
	beego.Get("/", func(ctx *context.Context) {

		// 用户数据的获取
		name := ctx.Input.Query("name")

		// 给用户响应数据
		ctx.Output.Context.WriteString(fmt.Sprintf("你输入的名字是：%s", name))
	})

	// 路由绑定函数,  路径，处理器函数（处理器）
	// 以POST方式请求，通过绑定函数处理
	beego.Post("/", func(ctx *context.Context) {
		// 用户数据的获取
		name := ctx.Input.Query("name")

		// 给用户响应数据
		ctx.Output.Context.WriteString(fmt.Sprintf("(POST)你输入的名字是：%s", name))
	})

	// Any函数任意请求都可以处理
	beego.Any("/any", func(ctx *context.Context) {
		// 用户数据的获取
		name := ctx.Input.Query("name")

		// 给用户响应数据,ctx.Input.Method() 获取请求方式
		ctx.Output.Context.WriteString(fmt.Sprintf("你输入的名字是：%s", name))
		ctx.Output.Context.WriteString(fmt.Sprintf("请求方式是：%s", ctx.Input.Method()))
	})

	// 启动beego
	beego.Run()
}
```

#### 2.2、动态路由（正则路由）

```go
// 导入beego的包
import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	// 正则路由
	// 匹配 /数字/ 格式的路由。并且把匹配到的值放入:id参数中
	beego.Get("/:id(\\d+)/", func(ctx *context.Context) {
		ctx.WriteString("匹配")
	})

	// 匹配任意的格式
	beego.Get("/any/:context/", func(ctx *context.Context) {
		ctx.WriteString("匹配context")
	})

	// 匹配一个文件，http://localhost:8080/file/xxx.aaa
	beego.Get("/file/*", func(ctx *context.Context) {
		ctx.WriteString("匹配file")
	})

	// 匹配ext，http://localhost:8080/ext/xxx.aaa
	beego.Get("/ext/*.*", func(ctx *context.Context) {
		ctx.WriteString("匹配ext")
	})

	// 启动beego
	beego.Run()
}
```

#### 2.3、url参数获取（少用，不建议用。不方便做权限控制）

```go
package main

// 导入beego的包
import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {

	// 正则路由
	// 匹配 /数字/ 格式的路由。并且把匹配到的值放入:id参数中
	beego.Get("/:id(\\d+)/", func(ctx *context.Context) {
		id := ctx.Input.Param(":id")
		ctx.WriteString(fmt.Sprintf("匹配的id：%s", id))
	})
	beego.Get("/:name(\\w+)/", func(ctx *context.Context) {
		name := ctx.Input.Param(":name")
		ctx.WriteString(fmt.Sprintf("匹配的name：%s", name))
	})

	// 匹配任意的格式
	beego.Get("/any/:context/", func(ctx *context.Context) {
		context := ctx.Input.Param(":id")
		ctx.WriteString(fmt.Sprintf("匹配的context：%s", context))
	})

	// 匹配一个文件，http://localhost:8080/file/xxx.aaa
	beego.Get("/file/*", func(ctx *context.Context) {
		// * 怎么表示呢？用splat
		splat := ctx.Input.Param(":splat")
		ctx.WriteString(fmt.Sprintf("匹配的file：%s", splat))
	})

	// 匹配ext，http://localhost:8080/ext/xxx.aaa
	beego.Get("/ext/*.*", func(ctx *context.Context) {
		// 匹配 . 前面的*用 path
		// 匹配 . 后面的用 ext
		path := ctx.Input.Param(":path")
		ext := ctx.Input.Param(":ext")
		ctx.WriteString(fmt.Sprintf("匹配ext前: %s,后：%s", path, ext))
	})

	// 启动beego
	beego.Run()
}
```



### 三、beego控制器

#### 3.1、Restful风格的控制器

```go
package main

// 导入beego的包
import (
	"github.com/astaxie/beego"
)

// 定义控制器
type HomeController struct {
	beego.Controller
}

/*
	1、Restful风格的控制器（API格式）
		a、一切皆资源 ----> url
		b、通过动作来表示对资源的操作类型  http method ---->  post 创建，Get 或者查询，Delete 删除，Put 更新
*/

// 1、可以处理Get请求
func (c *HomeController) Get() {
	c.Ctx.WriteString("Get")
}

// 2、可以处理Post请求
func (c *HomeController) Post() {
	c.Ctx.WriteString("Post")
}

// 3、可以处理Delete请求
func (c *HomeController) Delete() {
	c.Ctx.WriteString("Delete")
}

// 4、可以处理Put请求
func (c *HomeController) Put() {
	c.Ctx.WriteString("Put")
}

func main() {

	// 通过路由将URL跟控制器绑定
	beego.Router("/home", &HomeController{})

	beego.Run()
}
```

#### 3.2、自定义匹配控制器

```go
package main

import "github.com/astaxie/beego"

// 自定义控制器&路由规则
type TaskContorller struct {
	beego.Controller
}

// 添加任务
func (t *TaskContorller) Add() {
	t.Ctx.WriteString("Add")
}

// 查询任务
func (t *TaskContorller) Query() {
	t.Ctx.WriteString("Query")
}

// 删除任务
func (t *TaskContorller) Del() {
	t.Ctx.WriteString("Del")
}

// 修改任务
func (t *TaskContorller) Modify() {
	t.Ctx.WriteString("Modify")
}

func main() {

	// 自定义路由规则，用分号切割路由规则
	beego.Router("/task", &TaskContorller{}, "get,head:Query;post:Add;put:Modify")
	// 启动
	beego.Run()
}
```

#### 3.3、自动匹配路由，用的最多

```go
package main

import "github.com/astaxie/beego"

// 自定义控制器&路由规则
type TaskContorller struct {
	beego.Controller
}

// 添加任务
func (t *TaskContorller) Add() {
	t.Ctx.WriteString("Add")
}

// 查询任务
func (t *TaskContorller) Query() {
	t.Ctx.WriteString("Query")
}

// 删除任务
func (t *TaskContorller) Del() {
	t.Ctx.WriteString("Del")
}

// 修改任务
func (t *TaskContorller) Modify() {
	t.Ctx.WriteString("Modify")
}

func main() {

	// 自动路由，用的最多
	// url 控制 controller/action(方法)
	// add ----> Add方法
	beego.AutoRouter(&TaskContorller{})
	// 启动
	beego.Run()
}
```



### 四、获取url的数据

#### 4.1、获取数据

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
)

// 定义一个结构体，用来接收数据
type InputForm struct {
	Name     string
	Password string `form:"password"` // 定义标签就可以在url传递小写的参数了
}

type InputContorller struct {
	beego.Controller
}

// 从Get请求拿数据 （http://localhost:8080//inputcontorller/QueryParams/?Name=12）
func (c *InputContorller) QueryParams() {
	// 获取url的数据:
	/*
		方式一：不常用，最基础的方式
		c.Ctx.Request.ParseForm()
		fmt.Println(c.Ctx.Request.Form)

		方式二：
		fmt.Println(c.Ctx.Input.Query("name"))

		方式三：不常用，知道就行
		var name string
		c.Ctx.Input.Bind(&name, "name")
		fmt.Println(name)

		方式四：
		fmt.Println(c.GetString("name"))
	*/

	// 方式五：用的多。属性名跟URL传递的参数必须大小写一致（上面结构体的Name）。小写的话得定义标签，看上面结构体吧
	var form InputForm
	c.ParseForm(&form)
	fmt.Println(form)

	//
	c.Ctx.WriteString("")
}

// 从Post请求拿数据 （http://localhost:8080//inputcontorller/form/?Name=12）
func (c *InputContorller) Form() {
	/*
		方式一：
		c.Ctx.Request.ParseForm()
		fmt.Println(c.Ctx.Request.Form)
	*/

	// 方式二：
	c.Ctx.Request.ParseForm()
	fmt.Println(c.Ctx.Request.PostForm)
	// 方式三：
	fmt.Println(c.GetString("name"))

	// 方式四：用的多。属性名跟URL传递的参数必须大小写一致（上面结构体的Name）。小写的话得定义标签，看上面结构体吧
	var form InputForm
	c.ParseForm(&form)
	fmt.Println(form)

	//
	c.Ctx.WriteString("")
}

// 上传文件（http://localhost:8080//inputcontorller/file）
func (c *InputContorller) File() {
	// 方式一：使用Request 对象
	// c.Ctx.Request.FormFile("name")

	// 方式二：实际上就是封装了c.Ctx.Request.FormFile
	// c.GetFile("name")

	// 方式三：上传文件到当前路径upload目录下
	c.SaveToFile("img", "./upload/a.jpg")

	c.Ctx.WriteString("")
}

// json
func (c *InputContorller) Json() {
	// 一定要调用CopyBody，否则拿不到数据
	c.Ctx.Input.CopyBody(10 * 1024 * 1024) // 读取内容，后面会写到conf配置文件中
	var m map[string]interface{}

	json.Unmarshal(c.Ctx.Input.RequestBody, &m)     // 把json数据解析到 map中
	fmt.Println(string(c.Ctx.Input.RequestBody), m) // 解析数据
	c.Ctx.WriteString("")
}

func main() {

	beego.AutoRouter(&InputContorller{})
	beego.Run()
}
```

