package main

import (
	"fmt"
	"os"

	"github.com/Asker231/api_001.git/internal/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func main(){
	if err := godotenv.Load(".env");err != nil{
		fmt.Println(err.Error())
		return
	}
	db,err := gorm.Open(postgres.Open(os.Getenv("DNS")),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		return
	}

	db.AutoMigrate(&user.UserModel{})
}