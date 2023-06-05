package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"context"
	"time"
)

type (
	UpdateUserUseCase interface {
		Execute(context.Context, UpdateUserInput) (UpdateUserOutput, error)
	}

	UpdateUserInput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	UpdateUserPresenter interface {
		Output(domain.User) UpdateUserOutput
	}

	UpdateUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	updateUserUseCase struct {
		repository repository.UserRepository
		presenter  UpdateUserPresenter
		ctxTimeout time.Duration
	}
)

func NewUpdateUserUseCase(
	repository repository.UserRepository,
	presenter UpdateUserPresenter,
	t time.Duration,
) updateUserUseCase {
	return updateUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase updateUserUseCase) Execute(ctx context.Context, input UpdateUserInput) (UpdateUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	userDomain, err := domain.NewUser(input.Id, input.Name, input.Email)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}

	user, err := usecase.repository.UpdateItem(ctx, userDomain)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}

	return usecase.presenter.Output(user), nil
}
