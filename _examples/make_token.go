package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
	"wserver/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	tokenClaims := jwt.StandardClaims{
		Subject:   fmt.Sprintf("user%d", rand.Intn(9999)),
		ExpiresAt: time.Now().Unix() + 24*3600,
	}

	data, _ := json.Marshal(tokenClaims)
	fmt.Println(string(data))

	token, err := utils.MakeToken(tokenClaims)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Token: %s", token)
}
