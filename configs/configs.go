package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)


type(

	AppConfig struct{
		Auth AuthConfig
		DataBase DataBaseConfig
	}

	AuthConfig struct{

	}
	DataBaseConfig struct{
		DNS string
	}
)


func NewConfig()*AppConfig{
	if err := godotenv.Load(".env");err != nil{
		fmt.Println(err.Error())
		return nil
	}
	return &AppConfig{
		AuthConfig{},
		DataBaseConfig{
			DNS: os.Getenv("DNS"),
		},
	}
}