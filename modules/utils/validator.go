package utils

import (
	"encoding/json"
	"fmt"
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

func RouteValidator(method string, pattern string, handler func(http.ResponseWriter, *http.Request) (int, error), responseStruct interface{}) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println(responseStruct)
			// Validate the response payload following predefined schema
			err = validate.Struct(responseStruct)
			if err != nil {
				jsonEnc.Encode(Response{
					Message: err.Error(),
					Code:    http.StatusBadRequest,
				})
				return
			}

			code, err := handler(w, r)
			result = Response{
				Message: err,
				Code:    code,
				// Data:    data,
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
