package models

import (
	"errors"
	"practice/restapi/api-tsst/db"
	"practice/restapi/utils"
) 

type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
 hashedPassword,err:=utils.HashPassword(u.Password)
 if err!=nil{
	return err
 }
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}

	u.ID, err = result.LastInsertId()
	return err
}

func (u User)ValidateCredentails()error{
query:="SELECT password FROM users WHERE email=?"
row:=db.DB.QueryRow(query,u.Email)
var retrivewPassword string
 err:=row.Scan(&retrivewPassword)
 if err!=nil{
return err
 }
 passwrodIsValid:=utils.CheckPasswordHash(u.Password,retrivewPassword)
 if !passwrodIsValid{
	return  errors.New("credentials invalid")
 }
 return nil
}