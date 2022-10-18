package auth

type User struct {
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,min=8"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}
