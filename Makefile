GIT_VERSION := $(shell git describe --dirty --always --tags --long)
PACKAGE_NAME := $(shell go list -m -f '{{.Path}}')

EXECUTABLE := golang-api

JSON_SCHEMA_DIR := app/jsonschema
MIGRATION_DIR := app/migration
BIN_DIR := bin
GO_FILES := $(shell find . -type f -name '*.go')


.PHONY: all
all: ${BIN_DIR}/${EXECUTABLE}

.PHONY: wire
wire: app/service/wire_gen.go app/repository/wire_gen.go

# Bind json-schema files to the source
${JSON_SCHEMA_DIR}/bindata.go: ${JSON_SCHEMA_DIR} $(wildcard ${JSON_SCHEMA_DIR}/*.json)
	go-bindata -o $@ -prefix ${JSON_SCHEMA_DIR} -pkg jsonschema $(wildcard ${JSON_SCHEMA_DIR}/*.json)

# Bind SQL migration files to the source
${MIGRATION_DIR}/bindata.go: ${MIGRATION_DIR} $(wildcard ${MIGRATION_DIR}/*.sql)
	go-bindata -o $@ -prefix ${MIGRATION_DIR} -pkg migration $(wildcard ${MIGRATION_DIR}/*.sql)

# Resolve dependency injection
app/service/wire_gen.go: app/service/service_container.go
	go generate github.com/corentindeboisset/golang-api/app/service
app/repository/wire_gen.go: app/repository/repository_container.go
	go generate github.com/corentindeboisset/golang-api/app/repository

# Build the console
${BIN_DIR}/${EXECUTABLE}: ${GO_FILES} ${JSON_SCHEMA_DIR}/bindata.go ${MIGRATION_DIR}/bindata.go
	go build -ldflags "-X ${PACKAGE_NAME}/conf.Executable=${EXECUTABLE} -X ${PACKAGE_NAME}/conf.GitVersion=${GITVERSION}" -o $@
