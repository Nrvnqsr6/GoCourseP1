package model

import "github.com/golang-jwt/jwt/v4"

type UserRespone struct {
	user  *User
	token jwt.Token
}
