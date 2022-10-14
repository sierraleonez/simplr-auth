package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/utils"
)

func Route() {
	utils.RouteValidator("/", "GET", index)
	http.HandleFunc("/register", auth.Register)
}

type response struct {
	Status  int
	Message string
	Data    interface{}
}

func Middleware(pattern string, method string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(response{
				Status:  402,
				Message: "Unallowed Method",
				Data:    nil,
			})
			if err != nil {
				return
			}
			return
		}
	})
}

func register(w http.ResponseWriter, r *http.Request) {

}
func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "health check status: ok?")
}
