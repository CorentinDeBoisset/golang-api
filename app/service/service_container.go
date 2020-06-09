//+build wireinject

package service

import (
	"sync"
	"github.com/google/wire"
)

// Container contains all the services of the api.
// If you create a new service, be sure to add it in the Container and add its provider to the ServiceProviderSet
type Container struct {
	Connection	*Connection
	Logger		*Logger
	Migrator	*Migrator
}

var serviceProviderSet = wire.NewSet(
	wire.Struct(new(Container), "*"),
	ProvideConnection,
	ProvideLogger,
	ProvideMigrator,
)

var containerOnce sync.Once
var containerInstance *Container

func initializeContainer() (*Container, error) {
	wire.Build(serviceProviderSet)
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
