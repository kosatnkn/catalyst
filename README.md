# Catalyst
Go clean architecture RESTful API

[![Build Status](https://travis-ci.org/kosatnkn/catalyst.svg?branch=master)](https://travis-ci.org/kosatnkn/catalyst)
[![Coverage Status](https://coveralls.io/repos/github/kosatnkn/catalyst/badge.svg?branch=master)](https://coveralls.io/github/kosatnkn/catalyst?branch=master)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/catalyst)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/catalyst)
[![GoDoc](https://godoc.org/github.com/kosatnkn/catalyst?status.svg)](https://godoc.org/github.com/kosatnkn/catalyst)

## Features
- A server to handle web requests
- Configuration parser
- Container for dependency injection
- Router
- Controllers
- Error handler
- Logger
- Request mapper
- Response mapper and a Transformer
- Application Metrics


## Application Initialization Process

- Parse configurations
- Resolve container
- Initialize router
- Run server


## Request Response Cycle
```text
     + -------- +                + ------- +
     | RESPONSE |                | REQUEST |
     + -------- +                + ------- +
          /\                         ||
          ||                         \/
          ||                  + ------------ +  =>  + ---------- +
          ||                  |    Router    |      | Middleware |
          ||                  + ------------ +  <=  + ---------- +
          ||                             ||
          ||                             ||
     + --------------------------- +     ||
     | Transformer | Error Handler |     ||
     + --------------------------- +     ||
                                /\       ||
                                ||       \/
                            + -------------- +  =>  + --------- +
                            |   Controller   |      | Validator |
                            + -------------- +  <=  + --------- +
                                /\       ||
                                ||       \/
                            + -------------- +
                            |    Use Case    |
                            + -------------- +
                                /\       ||
                                ||       \/
              ______________________________________________
               + ------- +    + ---------- +    + ------- +
               | Adapter |    | Repository |    | Service |
               + ------- +    + ---------- +    + ------- +
                  /\  ||         /\    ||          /\  ||
                  ||  \/         ||    \/          ||  \/
               + ------- +    + ---------- +    + ------- +
               | Library |    |  Database  |    |   APIs  |
               + ------- +    + ---------- +    + ------- +
```


## View GoDoc Locally
```shell
    godoc -http=:6060 -v
```

Navigate to [http://localhost:6060/pkg/github.com/kosatnkn/catalyst/](http://localhost:6060/pkg/github.com/kosatnkn/catalyst/)


## Using Go mod

Go mod is used as the dependency management mechanism. Visit [here](https://github.com/golang/go/wiki/Modules) for more details.

- Use go mod in projects that are within the `GOPATH`
```bash
    export GO111MODULE=on
```

- View final versions that will be used in a build for all direct and indirect dependencies
```bash
    go list -m all
```
- View available minor and patch upgrades for all direct and indirect dependencies
```bash
    go list -u -m all
```
- Update all direct and indirect dependencies to latest minor or patch upgrades (pre-releases are ignored)
```bash
    go get -u or go get -u=patch
```
- Build or test all packages in the module when run from the module root directory
```bash
    go build ./... or go test ./...
```
- Prune any no-longer-needed dependencies from go.mod and add any dependencies needed for other combinations of OS, architecture, and build tags
```bash
    go mod tidy
```
- Optional step to create a vendor directory
```bash
    go mod vendor
```