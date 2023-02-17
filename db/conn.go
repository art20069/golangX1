package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB // ชนิดข้อมูลของตัวแปรตามค่าที่เราต้องการเก็บคือ db มีค่าเท่ากับ *gorm.DB

func ConnectDB() {

	dsn := os.Getenv("DATABASE_DSN")
	fmt.Println(dsn)
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{Logger: logger.Default.LogMode(logger.Info)},
	)
	if err != nil {
		log.Fatal("Connect conne;c;t to the database") // log ที่จะแสดงข้อความออกทาง teminal และปิดการทำงานให้ไม่ต้องมี return
	}
	Conn = db
}

// dsn := "user:1234567890@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
