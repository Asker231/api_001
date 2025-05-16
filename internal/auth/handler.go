package auth

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)


type AuthHandler struct{
   service *AuthService
}


func NewAuthHandler(app  *gin.RouterGroup,service *AuthService){
   authHandler := &AuthHandler{
      service: service,
   }

   app.POST("/register",authHandler.Register)
   app.POST("/login",authHandler.Login)
}



func(a *AuthHandler)Register(ctx *gin.Context){
   var user RegisterRequst
   if err := json.NewDecoder(ctx.Request.Body).Decode(&user);err != nil{
      return
   }
   returnType ,err := a.service.registerService(user.Name,user.Email,user.Password)
   if err != nil{
      ctx.Writer.Write([]byte(err.Error()))
      fmt.Println(err.Error())
      return 
   }
   data,_:= json.Marshal(&returnType)
   ctx.Writer.Write([]byte(data))

}

func(a *AuthHandler)Login(ctx *gin.Context){
   var user LoginRequst

   if err := json.NewDecoder(ctx.Request.Body).Decode(&user);err != nil{
      return
   }
   returnType,err := a.service.loginService(user.Email,user.Password)
   if err != nil{
      fmt.Println(err.Error())
      ctx.Writer.Write([]byte(err.Error()))
      return 
   }

   data,_:= json.Marshal(&returnType)
   ctx.Writer.Write([]byte(data))

}