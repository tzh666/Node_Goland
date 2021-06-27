package utils

import (
	"math/rand"
	"time"
)

// 生成随机字符串
func RandString(length int) string {

	lenters := "abcdefgijklmopqstuvwysABCDEFGIJKLMOPQRSWXYZ0123456789"
	count := len(lenters)
	chars := make([]byte, length)
	for i := 0; i < length; i++ {
		chars[i] = lenters[rand.Int()%count]
	}
	return string(chars)
}

// 设置种子
func init() {
	rand.Seed(time.Now().UnixNano())
}
