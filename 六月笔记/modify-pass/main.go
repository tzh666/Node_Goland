package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func GenerateFromPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		panic(err.Error())
	}
	return string(hash)
}

func CheckPassword(hash, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func main() {
	fmt.Println(GenerateFromPassword("123456"))
	fmt.Println(GenerateFromPassword("qqq111"))
	fmt.Println(CheckPassword("$2a$10$9y1eoX0l8B7ojwMWGnbGIOD6Uj52Tq7jueiooy1SQBiyWikasPXvC", "123456"))
	fmt.Println(CheckPassword(GenerateFromPassword("qqq111"), "qqq111"))
}
