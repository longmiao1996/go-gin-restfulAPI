package api

import (
	"encoding/base64"
	"fmt"
	"log"
	"miaolGo/drivers"
	"os"
	"strings"
	"time"
)

// AddArticle commit infos
func AddArticle(userName, nickName, title, article, imagebase64, image string) (flag bool) {
	imageName, imageFlag := UploadImage(userName, imagebase64, image)
	flag = false
	// 图片上传成功
	if imageFlag == true {
		db := drivers.Testsql()
		defer db.Close()
		stmt, err := db.Prepare("INSERT INTO article(name,nick_name,title,article,time,image_address) VALUES($1,$2,$3,$4,$5,$6)")
		if err != nil {
			log.Println(err)
			return
		}
		time := time.Now().Format("2006-01-02 15:04:05")
		_, err = stmt.Exec(userName, nickName, title, article, time, imageName)
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		return
	}
	return true

}

// UploadImage 图片上传
func UploadImage(userName, imageBase64, image string) (imageName string, flag bool) {
	// 前端传过来的base64格式需要修改
	imageBase64 = imageBase64[strings.IndexByte(imageBase64, ',')+1:]
	flag = true
	imageByte, err := base64.StdEncoding.DecodeString(imageBase64)
	if err != nil {
		fmt.Println(err)
		log.Println(err)
		flag = false
		return
	}
	time := time.Now().Format("2006-01-02 15:04:05")
	imageName = fmt.Sprintf("/media/image/%s%s%s.jpg", userName, time, image)
	file, err := os.OpenFile(imageName, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
		flag = false
		return
	}
	defer file.Close()
	file.Write(imageByte)

	return
}

// CheckFileExit 判断文件是否存在
func CheckFileExit(filePath string) (exit bool) {
	exit = true
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		exit = false
	}
	return
}
