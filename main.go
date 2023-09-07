package main

import (
	"go-litres/database"
	"go-litres/litres"
	"go-litres/logger"
	"go-litres/models"
	"sync"

	"os"

	"github.com/joho/godotenv"
)

func initDatabase() (db *database.DB) {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Создай файл с переменными окружениями .env!")
	}
	db = &database.DB{
		Host: os.Getenv("MYSQL_HOST"),
		User: os.Getenv("MYSQL_USER"),
		Name: os.Getenv("MYSQL_API_DATABASE"),
		Pass: os.Getenv("MYSQL_PASSWORD"),
		Port: os.Getenv("MYSQL_PORT"),
	}
	db.Connection = db.Connect()
	return
}

func main() {
	var lastId = 179378
	var wg sync.WaitGroup

	database := initDatabase()
	defer database.Close()
	for {
		var books []models.Book
		bookIds := database.Get100BooksID(lastId)
		if len(bookIds) == 0 {
			break
		}
		lastId = bookIds[len(bookIds)-1]
		wg.Add(len(bookIds))
		for _, id := range bookIds {
			go litres.PutBookToList(id, &wg, &books)
		}
		wg.Wait()
		if len(books) != 0 {
			database.SaveBooks(books)
		}
		logger.Log.Printf("Последняя книга: %d", lastId)
	}
	
}