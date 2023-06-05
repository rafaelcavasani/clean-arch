package usecase

import (
	"clean-arch/core/domain"
	"clean-arch/core/repository"
	"context"
	"time"
)

type (
	GetUserUseCase interface {
		Execute(context.Context, GetUserInput) (GetUserOutput, error)
	}

	GetUserInput struct {
		Id string `json:"id"`
	}

	GetUserPresenter interface {
		Output(domain.User) GetUserOutput
	}

	GetUserOutput struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	getUserUseCase struct {
		repository repository.UserRepository
		presenter  GetUserPresenter
		ctxTimeout time.Duration
	}
)

func NewGetUserUseCase(
	repository repository.UserRepository,
	presenter GetUserPresenter,
	t time.Duration,
) getUserUseCase {
	return getUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase getUserUseCase) Execute(ctx context.Context, input GetUserInput) (GetUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	user, err := usecase.repository.FindById(ctx, input.Id)
	if err != nil {
		return usecase.presenter.Output(domain.User{}), err
	}
	if user.Id == "" {
		return GetUserOutput{}, domain.ErrUserNotFound
	}

	return usecase.presenter.Output(user), nil
}
