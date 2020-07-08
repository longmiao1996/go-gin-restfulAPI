package models

import (
	"log"
	"miaolGo/drivers"
	"time"
)

type User struct {
	ID       int    `json:"id" form:"id" primaryKey:"true"`
	NAME     string `json:"name"`
	NICKNAME string `json:"nick_name"`
	PASSWORD string `json:"password"`
	EMAIL    string `json:"email"`
}

type LoginInfo struct {
	ID         int    `json:"id" form:"id" primarykey:"true"`
	CREATETIME string `json:"create_time"`
	USERNAME   string `json:"user_name"`
	USERPWD    string `json:"user_pwd"`
	STATUS     string `json:"status"`
}

func (model *User) CheckUser(name, pwd string) (bo bool, err error) {
	db := drivers.Testsql()
	defer db.Close()
	sqlStatement1 := `SELECT id FROM users WHERE name=$1 and password=$2;`
	var user User
	flag := true
	err = db.QueryRow(sqlStatement1, name, pwd).Scan(&user.ID)
	if err != nil {
		log.Println("login failed ")
		log.Println(err)
		flag = false
	}
	//把登录信息存入数据库login_info
	create_time := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("INSERT INTO login_info(create_time,user_name,user_pwd,status) VALUES($1,$2,$3,$4)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(create_time, name, pwd, flag)
	return flag, err
}
