package models

import (
	"database/sql"
	"log"
	"miaolGo/drivers"
)

var db *sql.DB = drivers.Testsql()

type Book struct {
	BOOKID    int    `json:"bookId" form:"bookId" primaryKey:"true"`
	BOOKPRICE string `json:"bookPrice,omitempty"`
	BOOKNAME  string `json:"bookName,omitempty"`
	BOOKAUTH  string `json:"bookAuth,omitempty"`
}

func (model *Book) GetUser(id int) (book Book, err error) {
	// find one record
	err = db.QueryRow("SELECT `bookAuth`, `bookPrice`, `bookName`, `bookID` FROM `books` WHERE `bookID` = ?;", id).Scan(&book.BOOKAUTH, &book.BOOKNAME, &book.BOOKNAME, &book.BOOKID)

	if err != nil {
		log.Println(err.Error())
		return
	}

	return
}
