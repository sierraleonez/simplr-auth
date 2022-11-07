package routes

import (
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/db"
	"simplr-auth/modules/utils"
)

type Test struct {
	// Ack string `validate:"required"`
}

var test = Test{}

func Route() {
	utils.RouteValidator("GET", "/", index, &test)
	utils.RouteValidator("POST", "/register", auth.Register, &auth.RequestForm)
	utils.RouteValidator("POST", "/login", auth.Login, &auth.LoginRequestForm)
}

func index(w http.ResponseWriter, r *http.Request) (int, interface{}, interface{}) {
	_, err := db.Connect()
	if err != nil {
		utils.Log(err.Error())
		return 501, err.Error(), nil
	}
	return 200, "health check status: ok?", nil
}
