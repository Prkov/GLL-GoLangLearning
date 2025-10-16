package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Мы пишем (Fprintln) в наш ответ (w) строку "Hello, World!"
	fmt.Fprintln(w, "Hello, World!")
}
func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "Date: %s\n", currentTime.Format("15:04:05"))
}
func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/time", timeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
