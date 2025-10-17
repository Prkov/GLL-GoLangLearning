package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Task struct {
	Id    int
	Title string
	IsOk  bool
}

var dataBase = []Task{
	{Id: 1, Title: "Основы GO", IsOk: false},
	{Id: 2, Title: "Создать первое API", IsOk: true},
	{Id: 3, Title: "Понять, что такое JSON", IsOk: false},
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Мы пишем (Fprintln) в наш ответ (w) строку "Hello, World!"
	fmt.Fprintln(w, "Hello, World!")
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now()
	fmt.Fprintf(w, "Date: %s\n", currentTime.Format("15:04:05"))
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var newTask Task
		err := json.NewDecoder(r.Body).Decode(&newTask)
		if err != nil {
			http.Error(w, "It's not a JSON", http.StatusInternalServerError)
			return
		} else {
			newTask.Id = dataBase[len(dataBase)-1].Id + 1
			dataBase = append(dataBase, newTask)
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newTask)
		}

	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		json.NewEncoder(w).Encode(dataBase)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

	}

}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
