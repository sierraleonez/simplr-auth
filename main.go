package main

import (
	"fmt"
	"net/http"

	"simplr-auth/modules/utils"
	"simplr-auth/routes"
)

func main() {
	routes.Route()
	env, err := utils.LoadConfig()
	if err != nil {
		utils.Log(err.Error())
		return
	}
	port := fmt.Sprintf(":%d", env.PORT)
	fmt.Printf("Starting web server at http://localhost%s/", port)
	http.ListenAndServe(port, nil)
}

type Message struct {
	Status string `json:"status"`
	Info   string `json:"info"`
}

// func handlePage(writer http.ResponseWriter, request *http.Request) {
// 	writer.Header().Set("Content-Type", "application/json")
// 	var message Message
// 	jwtToken, err := auth.GenerateJWT()

// 	message.Info = jwtToken
// 	message.Status = "Online"
// 	query.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
// 		var result User
// 		var email = "tester@yopmail.com"
// 		err = dbArg.
// 			QueryRow("select id, firstName, lastName from users where email = ?", email).
// 			Scan(&result.id, &result.firstName, &result.lastName)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return nil, err
// 		}
// 		fmt.Println(result)
// 		return result, nil
// 	})
// 	if err != nil {
// 		writer.WriteHeader(http.StatusUnauthorized)
// 		_, err := writer.Write([]byte("unable to generate token"))
// 		if err != nil {
// 			return
// 		}
// 		return
// 	}

// 	err = json.NewEncoder(writer).Encode(message)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// }
