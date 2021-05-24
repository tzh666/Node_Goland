## Golang之包

### 一、包的定义

#### 包的导入

要在代码中引用其他包的内容，需要使用`import`关键字导入使用的包。具体语法如下:

```
import "包的路径"
```

注意事项：

- import导入语句通常放在文件开头包声明语句的下面。
- 导入的包名需要使用双引号包裹起来。
- 包名是从`$GOPATH/src/`后开始计算的，使用`/`进行路径分隔。

#### 单行导入

单行导入的格式如下：

```
import "包1"
import "包2"
```

```go
https://blog.csdn.net/hyzx_9987/article/details/101059186
```



### 二、模块

```go


import (
    
    // 引用自己的 （module github.com/tzh666/testmod）是在go.mod 输入命令生成的
    // go mod init github.com/tzh666/testmod ，对应自己代码路径
    "github.com/tzh666/testmod/test"
    
    // 引用github官网的
	"github.com/astaxie/beego"
	"github.com/howeyc/gopass"
)

```

```go
// 其他的命令
go mod list
go mod tidy
等等
```

