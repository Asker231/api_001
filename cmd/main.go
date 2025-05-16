package main

import (
	"log"
	"net/http"

	"github.com/Asker231/api_001.git/configs"
	"github.com/Asker231/api_001.git/internal/auth"
	"github.com/Asker231/api_001.git/pkg/db"
	"github.com/gin-gonic/gin"
)

func main(){
	//init app
	app := gin.Default()
	v1  := app.Group("/auth")
	//config 
	cnf := configs.NewConfig()
	//db
	_ = db.NewConnectDb(&cnf.DataBase)

	//repo

	//service

	//handlers
	auth.NewAuthHandler(v1)



	server := http.Server{
		Addr: ":8080",
		Handler: app,
	}
	if err :=  server.ListenAndServe();err != nil{
		log.Fatal(err.Error())
		return 
	}
}