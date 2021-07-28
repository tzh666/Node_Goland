package utils

import (
	"io/ioutil"
	"os"
)

// 从文件中读取uuid
func ReadFile(path string) string {
	if ctx, err := ioutil.ReadFile(path); err != nil {
		return ""
	} else {
		return string(ctx)
	}
}

// 写入uuid
func WiteFile(path, txt string) {
	ioutil.WriteFile(path, []byte(txt), os.ModePerm)
}
