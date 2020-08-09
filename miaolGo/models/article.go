package models

// Article 记录用户发帖的信息
type Article struct {
	ID           int    `json:"id" form:"id" primaryKey:"true"`
	NAME         string `json:"name"`
	NICKNAME     string `json:"nick_name"`
	TITLE        string `json:"title"`
	ARTICLE      string `json:"article"`
	IMAGEADDRESS string `json:"image_address"`
	TIME         string `json:"time"`
}
