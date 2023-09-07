package database

import (
	"fmt"
	"go-litres/logger"
	"go-litres/models"
	"go-litres/tools"
)

func (d *DB) Get100BooksID(lastId int) (ids []int) {
	query := fmt.Sprintf(`SELECT id FROM book_xml_id WHERE id > %d ORDER BY id LIMIT 100`, lastId)
	rows, err := d.Connection.Query(query)
	tools.CheckErr(err)
	defer rows.Close()

	for rows.Next(){
		var id int
		err = rows.Scan(&id)
		tools.CheckErr(err)
		ids = append(ids, id)
	}
	return
}

func (d *DB) SaveManyBooks(books []models.Book) {
		groups := groupBooks(books)
		for _, group := range groups {
			d.SaveBooks(group)
		}
}

func (d *DB) SaveBooks(books []models.Book) {
	if len(books) == 0 { return }

	query, vals := createQueryForMultipleInsertBooksMYSQL(books)
	tx, _ := d.Connection.Begin()
	_, err := d.Connection.Exec(query, vals...)
	tools.CheckErr(err)
	tx.Commit()
	logger.Log.Printf("Успешно сохранили %d книг", len(books))
}


func groupBooks(books []models.Book) (groups [][]models.Book) {
	LIMIT := 2000
	for i := 0; i < len(books); i += LIMIT {
		group := books[i:]
		if len(group) >= 2000 {
			group = group[:LIMIT]
		}
		groups = append(groups, group)
	}
	return
}


func createQueryForMultipleInsertBooksMYSQL(books []models.Book) (query string, valArgs []interface{}) {
	query = `
		INSERT IGNORE INTO 
		book(id, name, description, language, price, old_price, currency, min_age, rating, year, image, url, pages, is_audio)
		VALUES `

	for _, book := range books {
		query += "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?),"
		valArgs = append(valArgs,  book.Id, book.Name, book.Description, book.Language, book.NewPrice, book.OldPrice, 
			book.Currency, book.MinAge, book.Rating, book.Year, book.Image, book.Url, book.Pages, book.IsAudio)
	}
	query = query[0:len(query)-1]
	return
}