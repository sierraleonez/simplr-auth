package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateJWT(data interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// Token expiration: 1 hour
	claims["exp"] = json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10))
	claims["authorized"] = true
	claims["user"] = data

	tokenString, err := token.SignedString([]byte("test"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header["Token"] != nil {
			token, err := jwt.Parse(request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				_, ok := token.Method.(*jwt.SigningMethodECDSA)
				if !ok {
					writer.WriteHeader((http.StatusUnauthorized))
					_, err := writer.Write([]byte("Invalid token"))
					if err != nil {
						return nil, err
					}
				}

				return "", nil
			})

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			if token.Valid {
				endpointHandler(writer, request)
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				_, err := writer.Write([]byte("You are unauthorized"))
				if err != nil {
					return
				}
			}

		} else {
			writer.WriteHeader(http.StatusUnauthorized)
			_, err := writer.Write([]byte("You're unauthorized due to no Token exist in your request header"))
			if err != nil {
				return
			}
		}
	})
}
