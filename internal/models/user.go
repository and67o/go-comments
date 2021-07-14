package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id int `json:"id"`
	Login string `json:"login"`
	Password string `json:"password"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}