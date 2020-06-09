GIT_VERSION := $(shell git describe --dirty --always --tags --long)
PACKAGE_NAME := $(shell go list -m -f '{{.Path}}')

EXECUTABLE := golang-api

JSON_SCHEMA_DIR := app/jsonschema
BIN_DIR := bin
GO_FILES := $(shell find . -type f -name '*.go')


.PHONY: all
all: ${BIN_DIR}/${EXECUTABLE}

# Bind json-schema files to the source
${JSON_SCHEMA_DIR}/bindata.go: $(wildcard ${JSON_SCHEMA_DIR}/*.json)
	go-bindata -o $@ -prefix ${JSON_SCHEMA_DIR} -pkg jsonschema $^

# Resolve dependency injection
app/service/wire_gen.go: app/service/service_container.go
	go generate github.com/corentindeboisset/golang-api/app/service
app/repository/wire_gen.go: app/repository/repository_container.go
	go generate github.com/corentindeboisset/golang-api/app/repository

# Build the console
${BIN_DIR}/${EXECUTABLE}: ${GO_FILES} ${JSON_SCHEMA_DIR}/bindata.go
	go build -ldflags "-X ${PACKAGE_NAME}/conf.Executable=${EXECUTABLE} -X ${PACKAGE_NAME}/conf.GitVersion=${GITVERSION}" -o $@
