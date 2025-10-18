package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // Наш "переводчик"
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
			http.Error(w, "It's not a JSON", http.StatusBadRequest)
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
	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable"
	// 2. Открываем "пул" соединений
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	defer db.Close() // Гарантирует, что соединение закроется при выходе из main

	// 3. Проверяем, что соединение живое
	err = db.Ping()
	if err != nil {
		log.Fatalf("База данных недоступна: %v", err)
	}

	fmt.Println("База данных успешно подключена!")

	// --- КОНЕЦ НОВОГО КОДА ---
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS tasks (
		id SERIAL PRIMARY KEY,
		title TEXT,
		is_ok BOOLEAN
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Не удалось создать таблицу: %v", err)
	}

	fmt.Println("Таблица 'tasks' готова к работе.")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/time", timeHandler)
	http.HandleFunc("/tasks", tasksHandler)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
