### Goland之结构体可见性

### 一、可见性

`首字母大写在包外才可以访问,属性也是同理u.brithdat！！！小写就不行`

```go
package main

import (
	"fmt"
	"kjx/user"
)

func main() {

	// 首字母大写在包外才可以访问,属性也是同理u.brithdat
	var u user.User
	//var a user.address

	fmt.Println(u.brithdat)
}
```

