package notes

type Notes struct {
	Id        string `validate:"required,number"`
	Title     string `validate:"required,alpha"`
	Content   string `validate:"alpha"`
	CreatedAt string `validate:"required,datetime"`
	UpdatedAt string `validate:"required,datetime"`
}

type CreateNoteRequest struct {
	Title   string `validate:"required,alpha"`
	Content string `validate:"required,alpha"`
}

type GetNoteByIdRequest struct {
	Id string `validate:"required" schema:"id"`
}

type InsertNoteRequest struct {
	Title   string `validate:"required,alpha"`
	Content string `validate:"alpha"`
}

type EditNoteByIdRequest struct {
	Id      string `validate:"required" schema:"id"`
	Title   string
	Content string
}

type DeleteNoteByIdRequest struct {
	Id string `validate:"required" schema:"id"`
}
