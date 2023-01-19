package auth

import (
	"database/sql"
	"errors"
	"net/http"
	model "simplr-auth/model/auth"
	customQuery "simplr-auth/modules/db/query"
	"simplr-auth/modules/utils"

	"golang.org/x/crypto/bcrypt"
)

var RequestForm model.RegisterRequest

func Register(w http.ResponseWriter, r *http.Request) (int, interface{}, interface{}) {
	// Check if user with registered email existed in DB
	emailExist, err := customQuery.QueryRow(
		func(dbArg *sql.DB) (res interface{}, err error) {
			var queryRes model.User
			err = dbArg.
				QueryRow("SELECT EXISTS(SELECT * FROM users WHERE email = ?)", RequestForm.Email).
				Scan(&queryRes.Email)
			if err != nil {
				return nil, err
			}
			return queryRes.Email, nil
		},
	)

	if err != nil {
		utils.Log(err.Error())
		return http.StatusInternalServerError, err.Error(), nil
	}

	if emailExist == "1" {
		return http.StatusBadRequest, errors.New("email already used by another user").Error(), nil
	} else {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(RequestForm.Password), 12)
		if err != nil {
			return http.StatusInternalServerError, err, nil
		}
		_, err = customQuery.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
			res, err = dbArg.Exec("INSERT INTO users (firstName, lastName, email, password) values(?, ?, ?, ?)", RequestForm.FirstName, RequestForm.LastName, RequestForm.Email, hashedPass)
			if err != nil {
				return nil, err
			}
			return
		})

		if err != nil {
			utils.Log(err.Error())
			return http.StatusInternalServerError, err.Error(), nil
		}
	}

	return http.StatusAccepted, "User registered", nil
}
