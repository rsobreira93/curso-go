package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/rsobreira93/curso-go/9-APIS/internal/dto"
	"github.com/rsobreira93/curso-go/9-APIS/internal/entity"
	"github.com/rsobreira93/curso-go/9-APIS/internal/infra/database"
)

type UserHandler struct {
	UserDb database.UserInterface
}

func NewUserHandler(db database.UserInterface) *UserHandler {
	return &UserHandler{
		UserDb: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserDTO

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := entity.NewUser(user.Name, user.Email, user.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.UserDb.Create(u)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
