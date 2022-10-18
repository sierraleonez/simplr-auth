package auth

import (
	"net/http"
	"simplr-auth/modules/utils"
)

type User struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

var RequestForm User

func Register(w http.ResponseWriter, r *http.Request) (int, error) {
	utils.Log(RequestForm)
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	err := r.ParseForm()
	if err != nil {
		utils.Log(err.Error())
		return http.StatusBadRequest, err
	}

	return http.StatusAccepted, nil
}
