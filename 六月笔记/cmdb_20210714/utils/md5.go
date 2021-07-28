package utils

import (
	"crypto/md5"
	"fmt"
)

// md5 解密
func Md5Text(text string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(text)))
}
