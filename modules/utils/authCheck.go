package utils

import (
	"errors"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

type ClaimStr struct {
	Id    string
	Email string
}

func AuthorizationCheck(w http.ResponseWriter, r *http.Request) (ClaimStr, error) {
	token := r.Header.Get("authorization")
	var claimsStruct ClaimStr
	env, err := LoadConfig()
	if err != nil {
		return claimsStruct, err
	}
	// - Check if token existed in request header
	if token != "" {
		// - Check if token is valid (by SHA256 signature and passcode provided)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(env.SECRET_TOKEN), nil
		})

		if err != nil {
			Log(err.Error())
			return claimsStruct, err
		}

		user := claims["user"].(map[string]interface{})
		parsedEmail := user["Email"].(string)
		parsedId := user["Id"].(string)
		claimsStruct = ClaimStr{
			Id:    parsedId,
			Email: parsedEmail,
		}

		Log("token is valid!`")
		return claimsStruct, nil

	}

	return claimsStruct, errors.New("no valid token in request")
}
