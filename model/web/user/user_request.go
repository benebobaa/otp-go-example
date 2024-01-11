package user

type UserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
