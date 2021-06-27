package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

/*
	md5 + 盐加密
	ext   要加密的内容
	salt  加的盐
*/
func Md5Salt(text string, salt string) string {
	if salt == "" {
		salt = RandString(8)
	}
	return fmt.Sprintf("%s:%x", salt, md5.Sum([]byte(fmt.Sprintf("%s:%s", text, salt))))
}

// 切片
func SplitMd5Salt(text string) (string, string) {
	// SplitN 分割两份
	nodes := strings.SplitN(text, ":", 2)
	if len(nodes) >= 2 {
		return nodes[0], nodes[1]
	} else {
		return "", nodes[0]
	}
}
