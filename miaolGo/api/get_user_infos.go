package api

import (
	"log"
	"miaolGo/drivers"
	"miaolGo/models"
)

//GetUserInfos 取得用户详细信息
func GetUserInfos(userName string) (nickName string, imageAddress string) {
	user := models.User{}
	db := drivers.Testsql()
	defer db.Close()
	sqlStatement1 := `SELECT id, nick_name, COALESCE(image_address,'') FROM users WHERE name=$1;`
	err := db.QueryRow(sqlStatement1, userName).Scan(&user.ID, &user.NICKNAME, &user.IMAGEADDRESS)
	if err != nil {
		log.Println(err)
		return
	}
	nickName = user.NICKNAME
	imageAddress = user.IMAGEADDRESS
	return

}
