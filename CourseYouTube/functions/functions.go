package main

import (
	"fmt"
)

func Greeting(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func main() {
	Greeting("Gopher")
}
