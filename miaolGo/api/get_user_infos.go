package api

import (
	"log"
	"miaolGo/drivers"
	"miaolGo/models"
)

// var struct User = models.User

func Get_user_infos(user_name string) (nick_name string, image_address string) {
	user := models.User{}
	db := drivers.Testsql()
	defer db.Close()
	sqlStatement1 := `SELECT id, nick_name, COALESCE(image_address,'') FROM users WHERE name=$1;`
	err := db.QueryRow(sqlStatement1, user_name).Scan(&user.ID, &user.NICKNAME, &user.IMAGEADDRESS)
	if err != nil {
		log.Println(err)
		return
	}
	nick_name = user.NICKNAME
	image_address = user.IMAGEADDRESS
	return

}
