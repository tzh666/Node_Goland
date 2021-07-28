package handler

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// 用户名密码
type AuthSecrets map[string]string

func Auth(handler http.Handler, secrets AuthSecrets) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		// 验证,获取 Basic YWFhOmFhYQ==
		secret := request.Header.Get("Authorization")
		if !isAuth(secret, secrets) {
			// 设置一个页面弹框,在页面输入用户名密码会被base64加密存储在Header中的Authorization里面(Authorization: Basic YWFhOmFhYQ==)
			// YWFhOmFhYQ== 这串东西解密以后就是username:password
			response.Header().Set("WWW-Authenticate", `Basic realm=""`)
			// 如果验证失败,返回401
			response.WriteHeader(401)
			return
		}
		//通过调用原生handler
		handler.ServeHTTP(response, request)
	})
}

// 判断
func isAuth(secret string, secrets AuthSecrets) bool {

	// 判断是否配置了用户名密码
	if secrets == nil {
		return true
	}

	// 解码Basic base64(username:password)  Authorization: Basic YWFhOmFhYQ==
	// secret 此时是"Basic YWFhOmFhYQ==" 我们只需要后面部分,所以用空格切割函数Fields
	nodes := strings.Fields(secret) // 返回一个切片  [Basic MTExOjIyMg==] nodes
	if len(nodes) != 2 {
		return false
	}

	// base64解密nodes[1]得到用户名密码username:password
	plaintext, err := base64.StdEncoding.DecodeString(nodes[1])
	if err != nil {
		return false
	}
	nodes = strings.SplitN((string(plaintext)), ":", 2)
	if len(nodes) != 2 {
		return false
	}

	hasher, ok := secrets[nodes[0]]
	fmt.Println(bcrypt.CompareHashAndPassword([]byte(hasher), []byte(nodes[1])))
	// hasher,password.如果hasher存在且密码正确(CompareHashAndPassword用于比对bcrypt哈希字符串和提供的密码明文文本是否匹配)
	return ok && bcrypt.CompareHashAndPassword([]byte(hasher), []byte(nodes[1])) == nil
}
