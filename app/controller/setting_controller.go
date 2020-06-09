package controller

import (
	"net/http"
	"github.com/corentindeboisset/golang-api/app/service"
	"github.com/corentindeboisset/golang-api/app/repository"
)

// SettingController handles routes related to Settings
type SettingController struct {
	ServiceContainer	*service.Container
	RepositoryContainer	*repository.Container
}

// InitializeSettingController returns an instance of a SettingController
func InitializeSettingController() (*SettingController, error) {
	serviceContainer, err := service.GetContainer()
	if err != nil {
		return nil, err
	}
	repositoryContainer, err := repository.GetContainer()
	if err != nil {
		return nil, err
	}

	return &SettingController{ServiceContainer: serviceContainer, RepositoryContainer: repositoryContainer}, nil
}

func (ctrl SettingController) SettingPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("userpage"))
}
