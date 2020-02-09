package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim ..., Token de usuario structura para la peticiones HTTP
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
