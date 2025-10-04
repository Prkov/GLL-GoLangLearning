package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Output struct {
	TimeSpent string
	Learned   string
	Notes     string
	Date      time.Time
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Ввод данных от пользователя
	fmt.Print("Time Spent (например, 1h 20m): ")
	timeSpent, _ := reader.ReadString('\n')
	timeSpent = strings.TrimSpace(timeSpent)

	fmt.Print("Learned (что изучил): ")
	learned, _ := reader.ReadString('\n')
	learned = strings.TrimSpace(learned)

	fmt.Print("Notes (заметки): ")
	notes, _ := reader.ReadString('\n')
	notes = strings.TrimSpace(notes)

	// Создаём структуру отчёта
	report := Output{
		TimeSpent: timeSpent,
		Learned:   learned,
		Notes:     notes,
		Date:      time.Now(),
	}

	// Формируем имя файла в формате YYYY-MM-DD.md
	filename := report.Date.Format("2006-01-02") + ".md"

	// Открываем файл для записи
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	// Пишем содержимое отчёта в Markdown формате
	content := fmt.Sprintf(
		"Date: %s  \nTime spent: %s  \nLearned: %s  \nNotes: %s\n",
		report.Date.Format("2006-01-02"),
		report.TimeSpent,
		report.Learned,
		report.Notes,
	)

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("\nОтчёт создан:", filename)
	fmt.Println("Его можно добавить в git и запушить на GitHub.")
}
