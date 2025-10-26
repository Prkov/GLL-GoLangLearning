package main

import (
	"fmt"
)

func Add(x, y int) int {
	return x + y
}
func Minus(x, y int) int {
	return x - y
}
func Multiply(x, y int) int {
	return x * y
}
func calculator(x, y int, action func(x, y int) int) int {
	return action(x, y)
}
func main() {
	fmt.Println(calculator(5, 4, Add))
}
