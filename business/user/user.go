package user

import "github.com/roby-aw/go-clean-architecture-hexagonal/utils"

type AuthLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type ResponseLogin struct {
	Email string      `json:"email"`
	Token utils.Token `json:"token"`
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (Register) timestamp(){
	
}