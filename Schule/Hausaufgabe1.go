package main

import (
	"fmt"
)

var counter int

func main() {
	for i := 0; i <= 50; i++ {
		counter += i
	}
	fmt.Println(counter)
}
