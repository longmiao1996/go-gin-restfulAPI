package api

import (
	"log"
	"miaolGo/drivers"
	"time"
)

// AddArticle commit infos
func AddArticle(userName, nickName, title, article, imageAddr string) (flag bool) {
	flag = true
	db := drivers.Testsql()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO article(name,nick_name,title,article,time,image_address) VALUES($1,$2,$3,$4,$5,$6)")
	if err != nil {
		log.Println(err)
		flag = false
		return
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(userName, nickName, title, article, time, imageAddr)
	if err != nil {
		log.Println(err)
		flag = false
		return
	}
	return

}
