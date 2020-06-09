//+build wireinject

package repository

import (
	"sync"
	"github.com/google/wire"
)

// Container contains all the repositories of the api.
// If you create a new repository, be sure to add it in the Container and add its provider to the repositoryProviderSet
type Container struct {
	SettingRepository	*SettingRepository
	UserRepository		*UserRepository
}

var repositoryProviderSet = wire.NewSet(
	wire.Struct(new(Container), "*"),
	ProvideSettingRepository,
	ProvideUserRepository,
)

var containerOnce sync.Once
var containerInstance *Container

func initializeContainer() (*Container, error) {
	wire.Build(repositoryProviderSet)
	return &Container{}, nil
}

// GetContainer returns always the same Container using a thread-safe singleton
func GetContainer() (*Container, error) {
	var err error = nil
	containerOnce.Do(func() {
		containerInstance, err = initializeContainer()
	})
	if err != nil {
		return nil, err
	}

	return containerInstance, nil
}
