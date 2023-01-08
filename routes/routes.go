package routes

import (
	"net/http"

	"simplr-auth/modules/auth"
	"simplr-auth/modules/db"
	"simplr-auth/modules/notes"
	"simplr-auth/modules/utils"
)

type Empty struct{}

var empty = Empty{}

func Route() {
	utils.RouteValidator("GET", "/", index, &empty)
	utils.RouteValidator("POST", "/register", auth.Register, &auth.RequestForm)
	utils.RouteValidator("POST", "/login", auth.Login, &auth.LoginRequestForm)

	// Notes
	utils.RouteValidator("GET", "/notes", notes.GetAllNotes, &empty)
	utils.RouteValidator("GET", "/notes/", notes.GetNoteById, &empty)
	utils.RouteValidator("POST", "/notes/create", notes.InsertNote, &notes.InsertRequestForm)
	utils.RouteValidator("PUT", "/notes/edit", notes.EditNoteById, &notes.EditNoteByIdRequest)
	utils.RouteValidator("DELETE", "/notes/delete", notes.DeleteNoteById, &notes.DeleteNoteByIdRequest)
}

func index(w http.ResponseWriter, r *http.Request) (int, interface{}, interface{}) {
	_, err := db.Connect()
	if err != nil {
		utils.Log(err.Error())
		return 501, err.Error(), nil
	}
	return 200, "health check status: ok?", nil
}
