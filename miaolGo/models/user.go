package models

import (
	"log"
	"miaolGo/drivers"
	"reflect"
	"time"
)

type User struct {
	ID           int    `json:"id" form:"id" primaryKey:"true"`
	NAME         string `json:"name"`
	NICKNAME     string `json:"nick_name"`
	PASSWORD     string `json:"password"`
	EMAIL        string `json:"email"`
	IMAGEADDRESS string `json:"image_address"`
}

type LoginInfo struct {
	ID         int    `json:"id" form:"id" primarykey:"true"`
	CREATETIME string `json:"create_time"`
	USERNAME   string `json:"user_name"`
	USERPWD    string `json:"user_pwd"`
	STATUS     string `json:"status"`
}

func (user *User) CheckUser(name, pwd, ip string) (flag bool, err error, nick_name string, image_address string) {
	db := drivers.Testsql()
	defer db.Close()
	sqlStatement1 := `SELECT id, nick_name, COALESCE(image_address,'') FROM users WHERE name=$1 and password=$2;`
	// var user User
	flag = true
	err = db.QueryRow(sqlStatement1, name, pwd).Scan(&user.ID, &user.NICKNAME, &user.IMAGEADDRESS)
	if err != nil {
		log.Println("login failed ")
		log.Println(err)
		flag = false
	}
	nick_name = user.NICKNAME
	image_address = user.IMAGEADDRESS
	//把登录信息存入数据库login_info
	create_time := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db.Prepare("INSERT INTO login_info(create_time,user_name,user_pwd,status,ip) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		log.Println(err)
	}
	_, err = stmt.Exec(create_time, name, pwd, flag, ip)
	return
}

func (user *User) AddUser(name, nick_name, password, email string) (err error, flag bool) {
	flag = false
	db := drivers.Testsql()
	defer db.Close()
	sqlStatement1 := `SELECT id FROM users WHERE name=$1;`
	// var user User
	err = db.QueryRow(sqlStatement1, name).Scan(&user.ID)
	if err != nil {
		log.Println(err)
		flag = true
	}
	//这个用户名没被注册过，可以继续注册
	if reflect.DeepEqual(*user, User{}) {
		stmt, err := db.Prepare("INSERT INTO users(name,nick_name,password,email,create_time) VALUES($1,$2,$3,$4,$5)")
		if err != nil {
			log.Println(err)
		}
		create_time := time.Now().Format("2006-01-02 15:04:05")
		_, err = stmt.Exec(name, nick_name, password, email, create_time)
	} else {
		return
	}
	return
}
