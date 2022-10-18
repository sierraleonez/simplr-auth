package auth

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	model "simplr-auth/model/auth"
	"simplr-auth/modules/db/query"
	"simplr-auth/modules/utils"

	"golang.org/x/crypto/bcrypt"
)

var LoginRequestForm model.LoginRequest

func Login(w http.ResponseWriter, r *http.Request) (int, interface{}, interface{}) {
	// select email and password from DB where email = request.email
	var user model.LoginRequest
	_, err := query.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
		err = dbArg.
			QueryRow("SELECT email, password FROM users WHERE email=?", LoginRequestForm.Email).
			Scan(&user.Email, &user.Password)

			// else, return error unidentified user
		if err != nil {
			utils.Log(err.Error())
			return nil, errors.New("user not found")
		}

		return user, nil
	})

	if err != nil {
		return http.StatusInternalServerError, err.Error(), nil
	}

	// compare request.password with DB.password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(LoginRequestForm.Password))
	if err != nil {
		return http.StatusUnauthorized, errors.New("password unmatched"), nil
	}

	// if true, create jwt
	token, err := utils.GenerateJWT(user)
	if err != nil {
		utils.Log(err)
		return http.StatusInternalServerError, errors.New("unable to create token"), nil
	}

	return http.StatusAccepted, fmt.Sprintf("Hello, %s", user.Email), token
}
