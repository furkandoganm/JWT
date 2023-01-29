package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("error different jwt method")
				}
				return EnvTokenKeyReader(), nil
			})
			EError(&w, "Parsing token!", err)

			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				AUser = claims["user"].(string)
				fmt.Printf("\n\n %v \n\n", claims["user"].(string))
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not outhorized")
		}
	})
}

var AUser string
