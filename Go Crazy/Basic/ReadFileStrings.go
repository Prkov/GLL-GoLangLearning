package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "logs"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error: %v \n", err)
	}
	defer file.Close()
	var counter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "ERROR") {
			fmt.Println(line)
		} else if strings.Contains(line, "INFO") {
			counter++
		}
	}
	fmt.Printf("Info repeated %v times \n", counter)

}
