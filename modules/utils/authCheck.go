package utils

import (
	// "encoding/json"
	// "fmt"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type ClaimStr struct {
	Id    string
	Email string
}

func AuthorizationCheck(w http.ResponseWriter, r *http.Request) (ClaimStr, error) {
	token := r.Header.Get("authorization")
	var claimsStruct ClaimStr
	// - Check if token existed in request header
	if token != "" {
		// - Check if token is valid (by SHA256 signature and passcode provided)
		claims := jwt.MapClaims{}
		_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte("test"), nil
		})

		if err != nil {
			Log(err.Error())
			return claimsStruct, err
		}
		Log(claims)
		user := claims["user"].(map[string]interface{})
		parsedEmail := user["Email"].(string)
		parsedId := user["Id"].(string)
		claimsStruct = ClaimStr{
			Id:    parsedId,
			Email: parsedEmail,
		}

		// Check if token had "exp" (expiration) field
		expiredDate, exist := claims["exp"]
		if !exist {
			Log("token invalid: doesnt have expiration field")
			return claimsStruct, err
		}

		// Assert expiration date type (expected to be int64)
		// reference https://stackoverflow.com/a/40946999/14593851
		tokenExp, ok := expiredDate.(int64)

		if !ok {
			Log("exp date not in valid format, expected: Int64")
		}

		// - Check token expiration
		// tokenExp := time.Unix(tx, 0)
		if tokenExp <= time.Now().Unix() {
			fmt.Println(tokenExp, time.Now().Unix(), expiredDate)
			Log("token expired")
			return claimsStruct, err
		}

		Log("token is valid!`")
		return claimsStruct, nil

	}
	return claimsStruct, errors.New("no valid token in request")
}
