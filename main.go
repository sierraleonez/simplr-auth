package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/db"
	"simplr-auth/modules/utils"
	"simplr-auth/routes"
)

type User struct {
	id        int
	firstName string
	lastName  string
}

func query() {
	email := "tester@yopmail.com"
	var result []User

	db, err := db.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, firstName, lastName from users where email = ?", email)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user = User{}
		var err = rows.Scan(&user.id, &user.firstName, &user.lastName)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.firstName)
	}
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Apa kabar?")
// }

func main() {
	routes.Route()
	// viper.SetConfigFile(".env")
	// viper.ReadInConfig()
	// fmt.Println(viper.Get("env"))
	// fmt.Println()
	utils.Trace_caller()
	http.HandleFunc("/a", handlePage)
	fmt.Println("Starting web server at http://localhost:8080/")
	http.ListenAndServe(":8000", nil)
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

func handlePage(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var message Message
	jwtToken, err := auth.GenerateJWT()

	message.Info = jwtToken
	message.Status = "Online"
	query()
	if err != nil {
		writer.WriteHeader(http.StatusUnauthorized)
		_, err := writer.Write([]byte("unable to generate token"))
		if err != nil {
			return
		}
		return
	}

	err = json.NewEncoder(writer).Encode(message)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
