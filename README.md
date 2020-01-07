# Catalyst

[![Build Status](https://travis-ci.org/kosatnkn/catalyst.svg?branch=master)](https://travis-ci.org/kosatnkn/catalyst)
[![Coverage Status](https://coveralls.io/repos/github/kosatnkn/catalyst/badge.svg?branch=master)](https://coveralls.io/github/kosatnkn/catalyst?branch=master)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/catalyst)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/catalyst)
[![GoDoc](https://godoc.org/github.com/kosatnkn/catalyst?status.svg)](https://godoc.org/github.com/kosatnkn/catalyst)

A REST API base that is written in **Go** using the **Clean Architecture** paradigm.

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

## Creating a New Project Using Catalyst

A new project can be created in one of two ways.

### Use Cauldron

The easiest way to create a project using `Catalyst` as the base is to use `Cauldron`. It is a small tool that enables you to set up a new project in no time.

More information about `Cauldron` can be found [here](https://github.com/kosatnkn/cauldron)

Clone and install `Cauldron`
```bash
    git clone https://github.com/kosatnkn/cauldron.git

    cd  cauldron

    go install
```

Create a new project
```bash
    $ cauldron -n=ProjectOne -ns=github.com/example [-t=v1.0.0]
```
> NOTE: 
> - -n Project name (ex: ProjectOne)
> - -ns Namespace for the project (ex: github.com/example)
> - -t Release version of Catalyst to be used. The latest version will be used if -t is not provided
> - -help or -h Show help message
 
Cauldron will do a git init on the newly created project but you will have to stage all the files in the project and do the first commit yourself.
```shell
    git add .

    git commit -m "first commit"
```

### Cloning

This is the work intensive approach.

Clone `Catalyst`
```bash
    git clone https://github.com/kosatnkn/catalyst.git <new_project_name>
```

Remove `.git`
```bash
    cd <new_project_name>

    rm -rf .git
```

Change import paths
> NOTE: Since `Catalyst` uses go mod the the newly created application will still work. But all the import paths would be as in `Catalyst` base project which is not what you will want.
- First change the module name in the `go.mod` file to a module name of your choice
- Then do a `Find & Replace` in the entire project to update all the import paths
- You may also need to change the splash text in `app/splash/styles.go`
- Now run and see whether the project compiles and run properly
- If so you can do a `git init` to the project

## The Sample Set
We have included a sample set of endpoints and their corresponding controller and domain logic by default.

This is to make it easier for you to follow through and understand how Catalyst handles the request response cycle for a given request.

The sample set will cover all basic CRUD operations that a REST API will normally need.

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
                            + -------------- +  =>  + -------------------- +
                            |   Controller   |      | Unpacker | Validator |
                            + -------------- +  <=  + -------------------- +
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

Use go mod in projects that are within the `GOPATH`
```bash
    export GO111MODULE=on
```

Initialize go mod
```bash
    go mod init github.com/my/repo
```

View final versions that will be used in a build for all direct and indirect dependencies
```bash
    go list -m all
```
View available minor and patch upgrades for all direct and indirect dependencies
```bash
    go list -u -m all
```
Update all direct and indirect dependencies to latest minor or patch upgrades (pre-releases are ignored)
```bash
    go get -u or go get -u=patch
```
Build or test all packages in the module when run from the module root directory
```bash
    go build ./... or go test ./...
```
Prune any no-longer-needed dependencies from go.mod and add any dependencies needed for other combinations of OS, architecture, and build tags
```bash
    go mod tidy
```
Optional step to create a vendor directory
```bash
    go mod vendor
```


## Docker

Catalyst provides a basic multistage Dockerfile so you have a starting point for creating Docker images.

```bash
    docker build -t <tag_name>:<tag_version> .
```

> NOTE: Do not forget the tailing `.` that indicates the current directory

**Example**
```bash
    docker build -t kosatnkn/catalyst:1.0.0 .
```

You can use it as follows
```bash
    docker run --name catalyst -p 3000:3000 -p 3001:3001 kosatnkn/catalyst:1.0.0
```

Do both in one go
```bash
    docker build -t kosatnkn/catalyst:1.0.0 . && docker run -it --rm --name catalyst -p 3000:3000 -p 3001:3001 kosatnkn/catalyst:1.0.0
```


## Wiki

Wiki pages on technical aspects of the project can be found [here](https://github.com/kosatnkn/catalyst/wiki)
