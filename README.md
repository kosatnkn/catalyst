# Catalyst
A `Clean Architecture` microservice template.

![catalyst_logo](./docs/img/catalyst_logo_256.svg)

[![CI](https://github.com/kosatnkn/catalyst/actions/workflows/ci.yml/badge.svg)](https://github.com/kosatnkn/catalyst/actions/workflows/ci.yml)
[![CodeQL](https://github.com/kosatnkn/catalyst/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/kosatnkn/catalyst/actions/workflows/codeql-analysis.yml)
[![Coverage Status](https://coveralls.io/repos/github/kosatnkn/catalyst/badge.svg?branch=master)](https://coveralls.io/github/kosatnkn/catalyst?branch=master)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/catalyst)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/catalyst)
[![Go Reference](https://pkg.go.dev/badge/github.com/kosatnkn/catalyst/v3.svg)](https://pkg.go.dev/github.com/kosatnkn/catalyst/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/kosatnkn/catalyst)](https://goreportcard.com/report/github.com/kosatnkn/catalyst)


## Introduction

For **version 3** of **Catalyst**, my main focus is for it to be simple, clean (obviously) and upgradable. Looking back, these are the very things that I struggled with both previous versions. Especially upgradability.

I have removed a substantial amount of code that I have passionately written for **Catalyst** over the years. In hindsight what I was doing was just reinventing the wheel again and again while there were better alternatives out there. More code I added for things like dynamic IoC container resolution, generalizing DB transactions and even logging and metric generation I was making Catalyst more and more opinionated.

With all of that bulk gone, Catalyst has become sort of a template. Now instead of saying what resources to use, Catalyst is saying where each type of resource should be put in to. You as the developer has the freedom to use whatever you feel fit for the implementation.

For an example you can use a simple struct implementation as the IoC container which is resolved manually. This is what I have done here for demo purposes (and also what I would actually do unless there is a very good reason to opt in something else) or you can drop in a heavy-duty IoC container. It is same for Loggers, Database adapters and what not.

I'm maintaining a separate GitHub repository at [kosatnkn/catalyst-pkgs](https://github.com/kosatnkn/catalyst-pkgs) to hold some packages that I use here. The logger is a wrapper around [rs/zerolog](https://github.com/rs/zerolog) and the configuration parser is a wrapper around [spf13/viper](https://github.com/spf13/viper). Feel free to swap them out for something better suites your needs.

## Architecture

There are many ways to organize a project to follow the **Clean Architecture** paradigm. This is how I organized **Catalyst**.

![Clean Architecture](./docs/img/clean_arch.drawio.svg)

When this architecture is mapped to the directory structure of **Catalyst**, it looks something like this.

![Clean Architecture Dir Mapping](./docs/img/clean_arch_dir_mapping.drawio.svg)

### Domain
The Domain contain all business logic executed by the microservice. It consist of three main parts; **Entities**, **Usecases** and the **Boundary**.

#### Entities
Entities define the data model for the domain. These are simple `Go` structs and are used within the domain as well as across the domain boundary to transfer data.

#### Usecases
Usecases contain all business logic. Other external dependencies that's needed by usecases (i.e. database resources) are injected in to these usecases by the use of dependency inversion.

#### Boundary
The Boundary marks the interface between the Domain and the orchestration layers. It contains contracts (`Go` interfaces) that facilitates dependency inversion.

### Orchestration
Orchestration contains **Infrastructure**, **Presentation** and **Persistence**.

#### Infrastructure
Infrastructure contains the lowest level resources needed by the microservice. Resources like configuration and the IoC container.

#### Presentation
Presentation contains all of outward facing interfaces. These are the communication channels to and from the microservice to the outside world. This is where you place your SESTful, GaphQL, GRPC, WebSocket servers. It is worth noting that you will not need to implement all of these in one microservice. It solely depends on specifics of your implementation.

Telemetry configs for metrics and traces can bee configured here as well. But with currently available options what I would do is to use a [eBPF](https://ebpf.io/) collector to collect telemetry. Unless you have to export custom metrics from your service this approach will provide sufficient information about your service.

#### Persistence
Persistence is used to keep data persistence related resources. Whether it is simple file writes, an RDBMS, an Object store or even an event sourcing system backed with a local store. The important thing to remember is that all implementation details should be encapsulated inside the Persistence layer. The domain that uses the persistence resources must not know nor care about how persistence is implemented in this layer. So saving to a static file should not be any different than to saving to a messaging backend when looked from the perspective of the Domain layer. All complexities replated to underlying persistence technologies should be contained within the Persistence layer.
