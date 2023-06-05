package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"encoding/json"
	"net/http"
)

type UpdateUserController struct {
	usecase usecase.UpdateUserUseCase
}

func NewUpdateUserController(usecase usecase.UpdateUserUseCase) UpdateUserController {
	return UpdateUserController{
		usecase: usecase,
	}
}

func (controller UpdateUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("userId")
	var input usecase.UpdateUserInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()
	input.Id = userId

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}
