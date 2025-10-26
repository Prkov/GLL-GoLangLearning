package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Подключение успешно")

}
