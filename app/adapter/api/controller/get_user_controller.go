package controller

import (
	"clean-arch/adapter/api/response"
	"clean-arch/core/usecase"
	"net/http"
)

type GetUserController struct {
	usecase usecase.GetUserUseCase
}

func NewGetUserController(usecase usecase.GetUserUseCase) GetUserController {
	return GetUserController{
		usecase: usecase,
	}
}

func (controller GetUserController) Execute(w http.ResponseWriter, r *http.Request) {
	var userId = r.URL.Query().Get("userId")
	var input = usecase.GetUserInput{Id: userId}

	output, err := controller.usecase.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusOK).Send(w)
}

func (controller GetUserController) handleErrors(w http.ResponseWriter, err error) {}
