package auth

import (
	"database/sql"
	"errors"
	"net/http"
	customQuery "simplr-auth/modules/db/query"
	"simplr-auth/modules/utils"
)

type User struct {
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}

var RequestForm User

func Register(w http.ResponseWriter, r *http.Request) (int, interface{}, interface{}) {
	// Check if user with registered email existed in DB
	emailExist, err := customQuery.QueryRow(
		func(dbArg *sql.DB) (res interface{}, err error) {
			var queryRes User
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
		insertRes, err := customQuery.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
			res, err = dbArg.Exec("INSERT INTO users (firstName, lastName, email, password) values(?, ?, ?, ?)", RequestForm.FirstName, RequestForm.LastName, RequestForm.Email, RequestForm.Password)
			if err != nil {
				return nil, err
			}
			return
		})

		if err != nil {
			utils.Log(err.Error())
			return http.StatusInternalServerError, err.Error(), nil
		}
		utils.Log(insertRes)
	}

	return http.StatusAccepted, "User registered", RequestForm
}
