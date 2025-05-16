package auth

import (
	"github.com/gin-gonic/gin"
)


type AuthHandler struct{
	
}


func NewAuthHandler(app  *gin.RouterGroup){
   authHandler := &AuthHandler{}

   app.POST("/register",authHandler.Register)
   app.POST("/login",authHandler.Login)
}



func(a *AuthHandler)Register(ctx *gin.Context){

}

func(a *AuthHandler)Login(ctx *gin.Context){

}