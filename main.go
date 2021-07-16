package main

import (
	"fmt"
	"github.com/charlloss/go-rest-api/init"
)

func main() {
	router := init.Routers()
	err := router.Run()
	if err != nil {
		fmt.Println("Server Run Error")
	}
}
