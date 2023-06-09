package config

import (
	"fmt"

	"project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	config := map[string]string{
		"DB_Username": "doadmin",
		"DB_Password": "AVNS_6OAAT5XU-qZFKRPpW5R",
		"DB_Port":     "25060",
		"DB_Host":     "testing-do-user-13944740-0.b.db.ondigitalocean.com",
		"DB_Name":     "training",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config["DB_Username"],
		config["DB_Password"],
		config["DB_Host"],
		config["DB_Port"],
		config["DB_Name"])

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{}, &models.Book{}, &models.Blog{})
}
