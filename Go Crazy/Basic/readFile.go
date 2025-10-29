package main

import (
	"fmt"
	"os"
)

func main() {
	content, err := os.ReadFile("Hello")

	if err != nil {
		fmt.Printf("Error: %v \n", err)
		return
	}
	fmt.Println(content)

}
