package main

import (
	"aivo-code-challenge/server"
	"fmt"
)

func main() {
	port := "8080"
	if err := server.New().Run(fmt.Sprintf(":%s", port)); err != nil {
		panic("PANIC on run server")
	}
}
