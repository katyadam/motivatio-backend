package main

import (
	"fmt"
	"motivatio-backend/api"
)

// RUN THIS: go run main.go
func main() {
	fmt.Println("Starting api...")
	api.Display()
	api.Db.Close()
}
