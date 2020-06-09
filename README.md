# Golang API boilerplate

API Boilerplate to start a project using golang

## Requirements

* GO >= 1.13

### Installation and first run

Install dependencies:

    go mod download

### Management

How to add a new service:

1. Add a file in `server/service`, with at least a Struct, and a Provider
2. Add the Service and the Struct in `server/service_container.go`
3. (optionnal) Run `go generate` to update the dependency injection. This should be run whenever `make` is run

Packages used:

+ router: chi & render
+ logger: zap
+ database: pq, sqlx & go-migrate
+ dependency injection: wire
+ static file injection: go-bindata
+ command line tool: cobra
+ config tool: viper
+ tests: testify
+ json schema: gojsonschema


id generation: xid ?
