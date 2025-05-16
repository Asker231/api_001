package user

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)


type RepositoryUser struct{
	Db *gorm.DB
}


func NewRepo(db *gorm.DB)*RepositoryUser{
	return &RepositoryUser{
		Db: db,
	}
}


func(r *RepositoryUser)CreateUser(user *UserModel)(*UserModel,error){
	userExist,_:= r.UserExist(user.Email)
	if userExist != nil{
		fmt.Println("UserExist")
		return nil,errors.New("Пользователь уже существует")
	}
	r.Db.Create(&user)

	return user,nil

	
}


func(r *RepositoryUser)UserExist(email string)(*UserModel,error){
	var user UserModel
	res := r.Db.First(&user,"email = ?",email)
	if res.Error != nil{
		fmt.Println(res.Error.Error())
		return nil,res.Error
	}
	return &user,nil
}