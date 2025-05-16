package db

import (
	"fmt"

	"github.com/Asker231/api_001.git/configs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


type ConnectDb struct{
	DataBase *gorm.DB
}


func NewConnectDb(cnf *configs.DataBaseConfig)*ConnectDb{

	db,err := gorm.Open(postgres.Open(cnf.DNS),&gorm.Config{})
	if err != nil{
		fmt.Println(err.Error())
		return nil
	}

	return &ConnectDb{
		DataBase: db,
	}
}