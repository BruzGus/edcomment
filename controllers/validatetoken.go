package controllers

import (
	"context"
	"net/http"

	"github.com/BruzGus/edcomment/commons"
	"github.com/BruzGus/edcomment/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

//ValidateToken ..., valida el token del cliente
func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	//respuesta al cliente
	var m models.Message
	token, err := request.ParseFromRequestWithClaims(
		r,                       //primer parametro es el request
		request.OAuth2Extractor, // tipo de extraccion
		&models.Claim{},         // destino donde se colocara la extraccion
		func(t *jwt.Token) (interface{}, error) {
			return commons.PublicKey, nil
		},
	)

	if err != nil {
		m.Code = http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			validationError := err.(*jwt.ValidationError)
			switch validationError.Errors {
			case jwt.ValidationErrorExpired:
				m.Message = "Su token a expirado"
				commons.DisplayMessage(w, m)
				return
			case jwt.ValidationErrorSignatureInvalid:
				m.Message = "La firma del cliente no coincide"
				commons.DisplayMessage(w, m)
				return
			default:
				m.Message = "Su toke no es valido"
				commons.DisplayMessage(w, m)
				return
			}
		}
	}

	if token.Valid {
		ctx := context.WithValue(
				r.Context(),
				"user",
				token.Claims.(*models.Claim).User)
			
		next(w, r.WithContext(ctx))
	} else {
		m.Code = http.StatusUnauthorized
		m.Message = "Su toke no es valido"
		commons.DisplayMessage(w, m)
	}

}
