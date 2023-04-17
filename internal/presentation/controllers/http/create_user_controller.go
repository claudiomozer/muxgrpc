package http

import (
	"encoding/json"
	"net/http"

	domain_user "github.com/claudiomozer/muxgrpc/internal/domain/user"
	controllers_protocols "github.com/claudiomozer/muxgrpc/internal/presentation/controllers/protocols"
)

type CreateUserController struct {
	create_user controllers_protocols.CreateUserUseCase
}

func NewCreateUserController(create_user controllers_protocols.CreateUserUseCase) *CreateUserController {
	return &CreateUserController{
		create_user: create_user,
	}
}

func (controller *CreateUserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var userDTO *domain_user.CreateUserUseCaseDTO

	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = controller.create_user.Create(*userDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
