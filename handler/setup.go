package handler

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Users struct {
	ID            int    `json:"id" gorm:"primaryKey;autoIncrement:true"`
	Name          string `json:"name"`
	Email         string `json:"email" gorm:"unique;not null;"`
	ContactNumber int    `json:"contactnumber"`
	Role          string `json:"role"`
	LibID         int    `json:"libid"`
}

var DB *gorm.DB

func Setup() {
	var err error
	dsn := "Nikhilsingh:password@tcp(127.0.0.1:3306)/NIKHIL?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connection not established proper")
	} else {
		fmt.Println("We are connected to the database ")
	}
	DB.AutoMigrate(&Users{})
}
