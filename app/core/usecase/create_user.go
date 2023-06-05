package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"context"
	"time"
)

type (
	CreateUserUseCase interface {
		Execute(context.Context, CreateUserInput) (CreateUserOutput, error)
	}

	CreateUserInput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	CreateUserPresenter interface {
		Output(domain.User) CreateUserOutput
	}

	CreateUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	createUserUseCase struct {
		repository repository.UserRepository
		presenter  CreateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateUserUseCase(
	repository repository.UserRepository,
	presenter CreateUserPresenter,
	t time.Duration,
) createUserUseCase {
	return createUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase createUserUseCase) Execute(ctx context.Context, input CreateUserInput) (CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	userDomain, err := domain.NewUser(input.Id, input.Name, input.Email)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}

	user, err := usecase.repository.PutItem(ctx, userDomain)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}

	return usecase.presenter.Output(user), nil
}
