package utils

import (
	"encoding/json"
	"net/http"
)

type Validators struct {
}

type response struct {
	Status  int
	Message string
	Data    interface{}
}

func RouteValidator(pattern string, method string, handler func(http.ResponseWriter, *http.Request)) {
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
