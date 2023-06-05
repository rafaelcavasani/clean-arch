package usecase

import (
	"clean-arch/core/repository"
	"context"
	"time"
)

type (
	DeleteUserUseCase interface {
		Execute(context.Context, DeleteUserInput) (DeleteUserOutput, error)
	}

	DeleteUserInput struct {
		Id string `json:"id"`
	}

	DeleteUserPresenter interface {
		Output(string) DeleteUserOutput
	}

	DeleteUserOutput struct {
		Id string `json:"id"`
	}

	deleteUserUseCase struct {
		repository repository.UserRepository
		presenter  DeleteUserPresenter
		ctxTimeout time.Duration
	}
)

func NewDeleteUserUseCase(
	repository repository.UserRepository,
	presenter DeleteUserPresenter,
	t time.Duration,
) deleteUserUseCase {
	return deleteUserUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (usecase deleteUserUseCase) Execute(ctx context.Context, input DeleteUserInput) (DeleteUserOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	err := usecase.repository.DeleteItem(ctx, input.Id)
	if err != nil {
		return usecase.presenter.Output(""), err
	}

	return usecase.presenter.Output(input.Id), nil
}
