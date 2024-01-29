package models

type User struct {
	ID       int    `json:"id" gorm:"primarykey;auto_increment:true;index"`
	Name     string `json:"name" gorm:"type:varchar(255)"`
	Password string `json:"password" gorm:"type:varchar(255)"`
	Email    string `json:"email" gorm:"type:varchar(255)"`
}

type Image struct {
	ID    int     `gorm:"primarykey;auto_increment:true;index"`
	Judul string  `gorm:"type:varchar(255)"`
	Foto  *string `gorm:"type:varchar(255)"`
}
