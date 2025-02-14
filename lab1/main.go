package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"lab1/calculators"
)

// Обробник для головної сторінки
func homePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Помилка завантаження сторінки", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Налаштування маршрутів
	http.HandleFunc("/", homePage)
	http.HandleFunc("/calculator1", calculators.CalculatorHandler1)
	http.HandleFunc("/calculator2", calculators.CalculatorHandler2)
	// Запуск сервера
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
