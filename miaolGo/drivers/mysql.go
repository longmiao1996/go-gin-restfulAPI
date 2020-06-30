package drivers

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"miaolGo/config"
)

func Testsql() *sql.DB {
	dbConfig := config.GetDbConfig()

	dbDSN := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbConfig["DB_HOST"],
		dbConfig["DB_PORT"],
		dbConfig["DB_USER"],
		dbConfig["DB_PWD"],
		dbConfig["DB_NAME"],
	)

	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		fmt.Println("error")
	}
	return db
}
