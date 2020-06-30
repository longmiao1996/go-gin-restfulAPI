package models

import (
	"database/sql"
	"log"
	"miaolGo/drivers"
)

var db *sql.DB = drivers.Testsql()

type Book struct {
	BOOKID    int    `json:"id" form:"id" primaryKey:"true"`
	BOOKPRICE string `json:"book_price,omitempty"`
	BOOKNAME  string `json:"book_name,omitempty"`
}

func (model *Book) GetUser(id int) (book Book, err error) {
	// find one record
	sqlStatement := `SELECT * FROM book WHERE id=$1;`
	err = db.QueryRow(sqlStatement, id).Scan(&book.BOOKPRICE, &book.BOOKNAME, &book.BOOKID)

	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}
