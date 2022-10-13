package routes

import (
	"fmt"
	"net/http"
)

func Route() {
	http.HandleFunc("/", index)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Apa kabar?")
}
