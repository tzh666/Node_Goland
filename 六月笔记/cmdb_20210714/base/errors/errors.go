package errors

import "github.com/astaxie/beego/validation"

type Errors struct {
	// 存放错误的名称，错误的信息
	errors map[string][]string
}

// 添加错误
func (e *Errors) AddAalidation(valid *validation.Validation) {
	if valid.HasErrors() {
		for key, errs := range valid.ErrorsMap {
			for _, err := range errs {
				e.Add(key, err.Message)
			}
		}
	}
}

// 添加错误
func (e *Errors) Add(key, err string) {
	// 判断某个key所对应的value是否存在, 第一个返回值是value(如果存在的话)
	if _, ok := e.errors[key]; !ok {
		e.errors[key] = make([]string, 0, 5)
	}
	// 添加错误
	e.errors[key] = append(e.errors[key], err)
}

// 获取所有错误
func (e *Errors) Errors() map[string][]string {
	// 直接返回错误信息
	return e.errors
}

// 通过name返回错误信息,map的value是一个切片，所以通过key获取到一个切片，返回就是一个切片
func (e *Errors) ErrorsByKey(key string) []string {
	// 直接返回错误信息
	return e.errors[key]
}

// 判断是否有错误，通过len(map)判断
func (e *Errors) HasErrors() bool {
	return len(e.errors) != 0 // ---> if len(e.errors)!=0 return 0
}

// 返回值是一个指针结构体
// 构造函数
func New() *Errors {
	// 返回一个结构体
	return &Errors{
		// 初始化errors
		errors: make(map[string][]string),
	}
}
