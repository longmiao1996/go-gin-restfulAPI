package models

import (
	"log"
)

type User struct {
	ID       int    `json:"id" form:"id" primaryKey:"true"`
	NAME     string `json:"name"`
	NICKNAME string `json:"nick_name"`
	PASSWORD string `json:"password"`
	EMAIL    string `json:"email"`
}

func (model *User) CheckUser(name, pwd string) (bo bool, err error) {
	sqlStatement1 := `SELECT * FROM fuck WHERE name=$1 and password=$2;`
	var user User
	err = db.QueryRow(sqlStatement1, name, pwd).Scan(&user.ID, &user.NAME, &user.NICKNAME, &user.PASSWORD, &user.EMAIL)
	if err != nil {
		log.Println("login failed ")
		log.Println(err)
		return false, err
	}

	return true, err
}
