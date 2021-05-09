package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	fmt.Println(string(hashedPassword))
}
