package main

import (
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"log"
)

func main() {
	// Используем переменную для текста, а не для HTML
	var pageText string
	cntxt, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	log.Println("Запускаю... Сейчас ты должен увидеть окно Chrome!")

	err := chromedp.Run(cntxt,
		chromedp.Navigate("https://www.google.com/"),
		chromedp.WaitVisible("#L2AGLb", chromedp.BySearch),
		chromedp.Click("#L2AGLb", chromedp.BySearch),
		chromedp.SendKeys(`[name="q"]`, "cat!\n"),
		chromedp.WaitVisible("body", chromedp.BySearch),
		chromedp.Text("body", &pageText, chromedp.ByQuery),
	)

	if err != nil {
		log.Fatal(err)
	}

	// time.Sleep(5 * time.Second) // Этот Sleep здесь не нужен, т.к. он был *после* выполнения Run

	// Выводим чистый текст
	fmt.Println(pageText)
	log.Println("Success")
}
