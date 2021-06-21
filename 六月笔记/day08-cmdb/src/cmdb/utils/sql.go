package utils

import (
	"fmt"
	"strings"
)

// Like sql查询字符串前后加 %
func Like1(q string) string {

	q = strings.TrimSpace(q)
	if q == "" {
		return ""
	}
	// 替换
	q = strings.Replace(q, "/", "//", -1)
	q = strings.Replace(q, "%", "/%", -1)
	q = strings.Replace(q, "_", "/_", -1)

	// Sprintf中使用 % 就再加一个%转义
	return fmt.Sprintf("%%%s%%", q)
}
