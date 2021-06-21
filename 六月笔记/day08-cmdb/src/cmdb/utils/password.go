package utils

import "golang.org/x/crypto/bcrypt"

// 密码加密
func GenerateFromPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

// 解密如果返回true，就是正确的
func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
