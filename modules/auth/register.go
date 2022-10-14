package auth

import (
	"encoding/json"
	"net/http"
)

type User struct {
	firstName string
	lastName  string
	email     string
}
type response struct {
	Status  int
	Message string
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
	} else {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(response{
			Status:  402,
			Message: "unallowed method",
		})
		if err != nil {
			return
		}
		// w.Write([]byte("Unallowed method"))
		return
	}
}
