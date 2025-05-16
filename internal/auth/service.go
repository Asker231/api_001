package auth

import (
	"github.com/Asker231/api_001.git/internal/user"
	"golang.org/x/crypto/bcrypt"
)



type AuthService struct{
	repo *user.RepositoryUser
}

func NewService(repo *user.RepositoryUser)*AuthService{

	return &AuthService{
		repo: repo,
	}
}

func(s *AuthService)registerService(name,email,password string)(*user.UserModel,error){

	  pass,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	  if err != nil{
		return nil,err
	 }
	 u := user.UserModel{
		Name: name,
		Email: email,
		Pssword: string(pass),
	}
	userCreated,err :=  s.repo.CreateUser(&u)
	if err != nil{
		return nil,err
	}

	return userCreated,nil
}

func(s *AuthService)loginService(email,password string)(*user.UserModel,error){
	userFinde ,err := s.repo.UserExist(email)
	if err != nil{
		return nil,err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(userFinde.Pssword),[]byte(password));err!= nil{
		return nil,err
	}
	return userFinde,nil
}