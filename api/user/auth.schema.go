package user

type RegisterUser struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=128"`
	LastName  string `json:"last_name" validate:"required,min=4,max=128"`
	Email     string `json:"email" validate:"required,email,min=4,max=128"`
	Password  string `json:"password" validate:"required,min=8,max=128"`
}

type LoginUser struct {
	Email    string `json:"email" validate:"required,email,min=4,max=128"`
	Password string `json:"password" validate:"required,min=8,max=128"`
}
