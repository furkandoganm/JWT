package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"projects/ProjectOneAPI_jwt_3/repository"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userName string
	uN, pwd, res := r.BasicAuth()
	BError(&w, "Chack your user name and password!", res)

	users, err := repository.GetUserJSON()
	EError(&w, "Getting users from json file!", err)

	for _, user := range users.User {
		if uN == user.UserName && pwd == user.Password {
			userName = user.UserName
		}
	}

	//BError(&w, "User was not find!", userName == "1")

	//fmt.Printf("\n\n %v \n\n", userName)
	token, err := GenerateJWT(userName)
	EError(&w, "Creating token!", err)

	fmt.Fprintf(w, token)
}

func GenerateJWT(uN string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = uN
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(EnvTokenKeyReader())

	return tokenString, err
}
