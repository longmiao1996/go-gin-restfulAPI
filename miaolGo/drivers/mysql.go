package drivers

import (
	"database/sql"
	"fmt"
	"miaolGo/config"
)


func Testsql() *sql.DB {
	dbConfig := config.GetDbConfig()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_NAME"],
		dbConfig["DB_CHARSET"],
	)

	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		fmt.Println("error")
	}
	return db
}
