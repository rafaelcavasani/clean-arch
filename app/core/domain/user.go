package domain

import (
	"clean-arch/infrastructure/logger"
	"errors"

	"github.com/go-playground/validator/v10"
)

var ErrUserNotFound = errors.New("user not found")

type User struct {
	Id    string `validate:"required,min=1,max=32"`
	Name  string `validate:"required,min=5,max=50"`
	Email string `validate:"required,email"`
}

func NewUser(id string, name string, email string) (User, error) {
	user := User{
		Id:    id,
		Name:  name,
		Email: email,
	}
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		logger.Errorf("Error validating user", err)
		return User{}, err
	}
	return user, nil
}
