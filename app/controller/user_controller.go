package controller

import (
	"net/http"
	"github.com/corentindeboisset/golang-api/app/service"
	"github.com/corentindeboisset/golang-api/app/repository"
)

// UserController handles routes related to users
type UserController struct {
	ServiceContainer	*service.Container
	RepositoryContainer	*repository.Container
}

// InitializeUserController returns an instance of a UserController
func InitializeUserController() (*UserController, error) {
	serviceContainer, err := service.GetContainer()
	if err != nil {
		return nil, err
	}
	repositoryContainer, err := repository.GetContainer()
	if err != nil {
		return nil, err
	}

	return &UserController{ServiceContainer: serviceContainer, RepositoryContainer: repositoryContainer}, nil
}

func (ctrl UserController) UserPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home"))
}
