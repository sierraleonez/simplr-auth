package utils

import "net/http"

// Handle CORS issue
// https://www.stackhawk.com/blog/golang-cors-guide-what-it-is-and-how-to-enable-it/
// 	Enable CORS to allow all request origin, all request method and all headers field
func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS, PUT, DELETE")

	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

}
