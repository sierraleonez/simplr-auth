package auth

type User struct {
	Id        string `validate:"required,alpha"`
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Email     string `validate:"required,email"`
}

type LoginRequest struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=8"`
}

type RegisterRequest struct {
	FirstName string `validate:"required,alpha"`
	LastName  string `validate:"required,alpha"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required,alpha"`
}

type LoginResponse struct {
	Token string
}
