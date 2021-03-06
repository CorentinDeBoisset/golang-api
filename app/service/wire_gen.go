// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package service

import (
	"github.com/google/wire"
	"sync"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from service_container.go:

func initializeContainer() (*Container, error) {
	logger, err := ProvideLogger()
	if err != nil {
		return nil, err
	}
	connection, err := ProvideConnection(logger)
	if err != nil {
		return nil, err
	}
	migrator, err := ProvideMigrator(logger, connection)
	if err != nil {
		return nil, err
	}
	container := &Container{
		Connection: connection,
		Logger:     logger,
		Migrator:   migrator,
	}
	return container, nil
}

// service_container.go:

// Container contains all the services of the api.
// If you create a new service, be sure to add it in the Container and add its provider to the ServiceProviderSet
type Container struct {
	Connection *Connection
	Logger     *Logger
	Migrator   *Migrator
}

var serviceProviderSet = wire.NewSet(wire.Struct(new(Container), "*"), ProvideConnection,
	ProvideLogger,
	ProvideMigrator,
)

var containerOnce sync.Once

var containerInstance *Container

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
