package routes

import (
	"fmt"
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/db"
	"simplr-auth/modules/utils"
)

type Empty struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Email     string `validate:"required,email"`
}

var test Empty = Empty{
	FirstName: "",
}

func Route() {
	// utils.RouteValidator("GET", "/", index, test)
	utils.RouteValidator("POST", "/register", auth.Register, &auth.RequestForm)
}

func index(w http.ResponseWriter, r *http.Request) (int, error) {
	_, err := db.Connect()
	if err != nil {
		utils.Log(err.Error())
		return 501, err
	}
	fmt.Fprintln(w, "health check status: ok?")
	return 200, nil
}
