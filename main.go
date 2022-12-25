package main

import (
	"ValidateComment/pkg/api"
	"ValidateComment/pkg/db"
	"log"
	"net/http"
)

func main() {
	// инициализация зависимостей приложения
	db, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Run mysql")

	api := api.New(db)

	log.Println("Run API")

	// запуск веб-сервера с API и приложением
	err = http.ListenAndServe(":83", api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
