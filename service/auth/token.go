package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJWT(secret []byte, userId string, email string) (string,error) {
	expiration := time.Hour * time.Duration(1)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":userId,
		"email":email,
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString,err:=token.SignedString(secret)
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}

func VerifyToken(tokenString string,secret []byte,) error{
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secret, nil
   })
  
   if err != nil {
      return err
   }
  
   if !token.Valid {
      return fmt.Errorf("invalid token")
   }
  
   return nil
}
