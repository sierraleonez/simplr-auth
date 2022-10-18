package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type response struct {
	Status  int
	Message string
	Data    interface{}
}

type Response struct {
	Message interface{}
	Code    int
	Data    interface{}
}

// Validate method and payload
//
// [method] - method of the request (GET, POST, PUT, DELETE, etc)
//
// [pattern] - URI path of the request (/register, /login)
//
// [handler] - handler function to be executed
//
// [responseStruct] - pointer to instance of the request struct
func RouteValidator(method string, pattern string, handler func(http.ResponseWriter, *http.Request) (int, interface{}, interface{}), responseStruct interface{}) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

		var result Response
		jsonEnc := json.NewEncoder(w)
		decoder := schema.NewDecoder()
		validate := validator.New()
		if r.Method == method {

			// Map the response to struct following predefined schema
			r.ParseForm()
			err := decoder.Decode(responseStruct, r.Form)
			if err != nil {
				jsonEnc.Encode(Response{
					Message: err.Error(),
					Code:    http.StatusBadRequest,
				})
				return
			}

			// Validate the response payload following predefined schema
			err = validate.Struct(responseStruct)
			if err != nil {
				jsonEnc.Encode(Response{
					Message: err.Error(),
					Code:    http.StatusBadRequest,
				})
				return
			}

			code, message, data := handler(w, r)
			result = Response{
				Message: message,
				Code:    code,
				Data:    data,
			}
			jsonEnc.Encode(result)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			err := jsonEnc.Encode(response{
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
