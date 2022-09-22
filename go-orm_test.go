package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

const DB_USERNAME = "root"
const DB_PASSWORD = ""
const DB_NAME = "my_db"
const DB_HOST = "localhost"
const DB_PORT = "3306"

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func (p *Product) ToString() string {
	return fmt.Sprintf("<Product Code=%v, Price=%v>", p.Code, p.Price)
}

func connectDB() *gorm.DB {
	var err error
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("Error connecting to database : error=%v", err)
		return nil
	}
	fmt.Printf("Connected to db %v", dsn)

	return db
}

func Test_gorm1(t *testing.T) {
	log.Printf("ok")
	db := connectDB()
	db.AutoMigrate(&Product{})

	var p Product
	db.Where("Price > ?", 10).Find(&p)
	fmt.Printf("data: %v \n", p.ToString())
}
