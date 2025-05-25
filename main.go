// /*
// Copyright © 2025 NAME HERE <EMAIL ADDRESS>
// */
package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/yashpal2104/json-viewer-cli/cmd"
	// "github.com/yashpal2104/json-viewer-cli/data"
)

func main() {
	cmd.Execute()
	dummyJWT()
}

func dummyJWT() {
	token := jwt.New(jwt.SigningMethodHS256)
	tokenString, _ := token.SignedString([]byte("secret"))
	fmt.Println("Dummy token:", tokenString)
}
