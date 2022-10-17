package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type response struct {
	Status  int
	Message string
	Data    interface{}
}
type request struct {
	firstName string `validate:"required"`
	lastName  string `validate:"required"`
	email     string `validate:"required,email"`
}

func ValidateStruct(arg interface{}) error {
	validate := validator.New()
	err := validate.Struct(arg)

	if err != nil {
		return err
	}
	return nil
}

func RouteValidator(method string, pattern string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			err := json.NewEncoder(w).Encode(response{
				Status:  405,
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
