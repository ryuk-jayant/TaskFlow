package auth

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashFromPassword(password string)(string,error){
	hash, error:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if error!=nil{
		return "",fmt.Errorf("Error in Encryption!")
	}
	return string(hash),nil
}

func ComparePasswords(hashPassword string,password []byte)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashPassword),password)
	return err==nil
}