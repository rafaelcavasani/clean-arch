package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"encoding/json"
	"net/http"
)

type CreateUserController struct {
	usecase usecase.CreateUserUseCase
}

func NewCreateUserController(usecase usecase.CreateUserUseCase) CreateUserController {
	return CreateUserController{
		usecase: usecase,
	}
}

func (controller CreateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
