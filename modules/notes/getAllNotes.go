package notes

import (
	"database/sql"
	"errors"
	"net/http"
	model "simplr-auth/model/notes"
	customQuery "simplr-auth/modules/db/query"
	"simplr-auth/modules/utils"
	"strings"
)

func GetAllNotes(w http.ResponseWriter, r *http.Request) (status int, message interface{}, data interface{}) {
	var notes []model.Notes
	_, err := customQuery.QueryRow(
		func(dbArg *sql.DB) (res interface{}, err error) {
			var queryRes model.Notes
			result, err := dbArg.
				Query("SELECT * FROM notes")

			if err != nil {
				return nil, err
			}

			for result.Next() {
				note := model.Notes{}
				err = result.Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

				if err != nil {
					utils.Log(err.Error())
					return nil, err
				}

				notes = append(notes, note)
			}
			return queryRes, nil
		},
	)

	if err != nil {
		utils.Log(err.Error())
		return http.StatusInternalServerError, err.Error(), nil
	}

	return http.StatusAccepted, "Success getting notes", notes
}

var InsertRequestForm model.InsertNoteRequest

func InsertNote(w http.ResponseWriter, r *http.Request) (status int, message interface{}, data interface{}) {
	_, err := customQuery.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
		res, err = dbArg.Exec("INSERT INTO notes (title, content) values(?, ?)", InsertRequestForm.Title, InsertRequestForm.Title)
		if err != nil {
			utils.Log(err.Error())
			return nil, err
		}
		return
	})

	if err != nil {
		utils.Log(err.Error())
		return http.StatusInternalServerError, err.Error(), nil
	}

	return http.StatusOK, "success insert to DB", nil
}

var EditNoteByIdRequest model.EditNoteByIdRequest

func EditNoteById(w http.ResponseWriter, r *http.Request) (status int, message interface{}, data interface{}) {
	noteId := EditNoteByIdRequest.Id
	noteExist, err := CheckNoteExistById(noteId)

	if err != nil {
		utils.Log(err.Error())
		return http.StatusInternalServerError, err.Error(), nil
	}

	if noteExist == "1" {
		_, err := customQuery.QueryRow(
			func(dbArg *sql.DB) (res interface{}, err error) {
				var queryRes string
				res, err = dbArg.
					Exec("UPDATE notes SET title = ?, content = ? WHERE id = ?", EditNoteByIdRequest.Title, EditNoteByIdRequest.Content, noteId)
				if err != nil {
					utils.Log(res)
					utils.Log(err.Error())
					return nil, err
				}

				return queryRes, nil
			},
		)

		if err != nil {
			utils.Log(err.Error())
			return http.StatusInternalServerError, err.Error(), nil
		}
		return http.StatusOK, "exist", noteId
	} else {
		return http.StatusBadRequest, "Id not found", nil
	}
}

var DeleteNoteByIdRequest model.DeleteNoteByIdRequest

func DeleteNoteById(w http.ResponseWriter, r *http.Request) (status int, message interface{}, data interface{}) {
	noteId := DeleteNoteByIdRequest.Id
	isExist, err := CheckNoteExistById(noteId)

	if err != nil {
		utils.Log(err.Error())
		return http.StatusBadRequest, err.Error(), nil
	}

	if isExist == "1" {
		_, err = customQuery.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
			res, err = dbArg.Exec("DELETE FROM notes WHERE id = ?", noteId)

			if err != nil {
				utils.Log(err.Error())
				return nil, err
			}

			return res, nil
		})

		if err != nil {
			utils.Log(err.Error())
			return http.StatusInternalServerError, err.Error(), nil
		}

		return http.StatusOK, "delete success", nil
	} else {
		return http.StatusBadRequest, "id not found", nil
	}
}

// var GetNoteByIdRequest model.GetNoteByIdRequest

func GetNoteById(w http.ResponseWriter, r *http.Request) (status int, message interface{}, data interface{}) {
	noteId, err := SlugExtractor("/notes/", r)
	var note model.Notes
	if err != nil {
		utils.Log(noteId)
		return http.StatusBadRequest, "no slug", nil
	}

	_, err = customQuery.QueryRow(func(dbArg *sql.DB) (res interface{}, err error) {
		err = dbArg.
			QueryRow("SELECT * FROM notes WHERE id = ?", noteId).
			Scan(&note.Id, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt)

		return nil, err
	})

	if err != nil {
		utils.Log(err.Error())
		return http.StatusInternalServerError, err.Error(), nil
	}

	return http.StatusOK, "success", note

}

// Utils

func SlugExtractor(prefix string, r *http.Request) (slug string, err error) {
	url := r.URL.Path
	if strings.HasPrefix(url, prefix) {
		slug = url[len(prefix):]
		return slug, nil
	} else {
		return "", errors.New("URL contain no slug")
	}

}

func CheckNoteExistById(noteId string) (interface{}, error) {
	noteExist, err := customQuery.QueryRow(
		func(dbArg *sql.DB) (res interface{}, err error) {
			var queryRes string

			err = dbArg.
				QueryRow("SELECT EXISTS(SELECT * FROM notes WHERE id = ?)", noteId).
				Scan(&queryRes)

			if err != nil {
				utils.Log(err.Error())
				return nil, err
			}

			return queryRes, nil
		},
	)
	return noteExist, err
}
