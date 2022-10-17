package routes

import (
	"fmt"
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/utils"
)

func Route() {
	utils.RouteValidator("GET", "/", index)
	utils.RouteValidator("GET", "/register", auth.Register)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "health check status: ok?")
}
