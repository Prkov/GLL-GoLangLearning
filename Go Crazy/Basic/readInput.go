package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Print("\nYour input: ")
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Written Text: %v \n", line)
	}
}
