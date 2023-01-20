package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/schema"
)

type Response struct {
	Message interface{}
	Code    int
	Data    interface{}
}

// Validate method and payload
//
// GET method only able to make use of URL Query parameter and unable to make use of form value parameter,
// Other method should have their payload encoded in x-www-form-urlencoded
//
// [method] - method of the request (GET, POST, PUT, DELETE, etc)
//
// [pattern] - URI path of the request (/register, /login)
//
// [handler] - handler function to be executed
//
// [requestStruct] - pointer to instance of the request struct
func RouteValidator(method string, pattern string, isProtected bool, handler func(http.ResponseWriter, *http.Request) (int, interface{}, interface{}), requestStruct interface{}) {
	http.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		var result Response
		decoder := schema.NewDecoder()
		validate := validator.New()

		// Handle CORS, allow all origin
		EnableCors(&w)

		// Check token in request if route is protected
		if isProtected {
			_, err := AuthorizationCheck(w, r)
			if err != nil {
				CreateResponse(w, Response{
					Message: err.Error(),
					Code:    http.StatusUnauthorized,
				})
				return
			}
		}

		w.Header().Set("Content-Type", "application/x-www-form-urlencoded")

		if r.Method == method {
			// Map the request body to struct following predefined schema
			r.ParseForm()
			err := decoder.Decode(requestStruct, r.Form)
			if err != nil {
				CreateResponse(w, Response{
					Message: err.Error(),
					Code:    http.StatusBadRequest,
				})
				return
			}
			Log(r.Form)
			// Validate the response payload following predefined schema
			err = validate.Struct(requestStruct)
			if err != nil {
				CreateResponse(w, Response{
					Message: err.Error(),
					Code:    http.StatusBadRequest,
				})
				return
			}

			// Execute handler method and retrieve the result (if exist)
			code, message, data := handler(w, r)
			result = Response{
				Message: message,
				Code:    code,
				Data:    data,
			}
			CreateResponse(w, result)

		} else {

			// Handle CORS preflight request
			// (https://developer.mozilla.org/en-US/docs/Glossary/Preflight_request#:~:text=A%20CORS%20preflight%20request%20is,Headers%20%2C%20and%20the%20Origin%20header.)
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			w.WriteHeader(http.StatusBadRequest)
			err := CreateResponse(w, Response{
				Code:    http.StatusMethodNotAllowed,
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

func CreateResponse(w http.ResponseWriter, res Response) error {
	Log(res)
	jsonEnc := json.NewEncoder(w)
	w.WriteHeader(res.Code)
	err := jsonEnc.Encode(res)
	return err
}
