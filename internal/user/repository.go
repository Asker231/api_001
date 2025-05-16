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

	userExist ,err := r.UserExist(user.Email)

	if err != nil{
		fmt.Println(err)
		return nil,err
	}

	if userExist != nil{
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