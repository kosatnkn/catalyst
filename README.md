# Catalyst
A `Clean Architecture` microservice template written in `Go`.

![catalyst_logo](./docs/img/catalyst_logo.svg)

[![CI](https://github.com/kosatnkn/catalyst/actions/workflows/ci.yml/badge.svg)](https://github.com/kosatnkn/catalyst/actions/workflows/ci.yml)
[![CodeQL](https://github.com/kosatnkn/catalyst/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/kosatnkn/catalyst/actions/workflows/codeql-analysis.yml)
[![Coverage Status](https://coveralls.io/repos/github/kosatnkn/catalyst/badge.svg?branch=master)](https://coveralls.io/github/kosatnkn/catalyst?branch=master)
![Open Issues](https://img.shields.io/github/issues/kosatnkn/catalyst)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/catalyst)
[![Go Reference](https://pkg.go.dev/badge/github.com/kosatnkn/catalyst/v3.svg)](https://pkg.go.dev/github.com/kosatnkn/catalyst/v3)
[![Go Report Card](https://goreportcard.com/badge/github.com/kosatnkn/catalyst)](https://goreportcard.com/report/github.com/kosatnkn/catalyst)

## 1. Introduction

For **version 3** of **Catalyst**, my main focus is to make it simple, clean and upgradable. Looking back, these are the very things I struggled with in both previous versions. Especially upgradability.

I have removed a substantial amount of code that I had passionately written for **Catalyst** over the years. In hindsight, I realized that I was just reinventing the wheel again and again, while better alternatives already existed. The more code I added for things like dynamic IoC container resolution, generalized DB transactions, and even logging and metric generation, the more opinionated **Catalyst** became.

With all that bulk gone, **Catalyst** has become more of a template. Now, instead of dictating which resources to use, **Catalyst** simply defines where each type of resource should go. As the developer, you have the freedom to choose whatever you feel best fits your implementation.

For example, you can use a simple struct implementation as the IoC container and resolve it manually. This is what I’ve done here for demo purposes (and also what I would actually do unless there’s a very good reason to opt for something else). Alternatively, you can plug in a heavy-duty IoC container. The same applies to loggers, database adapters, and so on.

I’m maintaining a separate GitHub repository at [kosatnkn/catalyst-pkgs](https://github.com/kosatnkn/catalyst-pkgs) to host some of the packages I use here. The logger is a wrapper around [rs/zerolog](https://github.com/rs/zerolog), and the configuration parser is a wrapper around [spf13/viper](https://github.com/spf13/viper). Feel free to swap them out for whatever better suits your needs.

## 2. Architecture

There are many ways to organize a project that follows the **Clean Architecture** paradigm. This is how I’ve organized **Catalyst**.

![Clean Architecture](./docs/img/clean_arch.drawio.svg)

When this architecture is mapped to the directory structure of **Catalyst**, it looks like this.

![Clean Architecture Dir Mapping](./docs/img/clean_arch_dir_mapping.drawio.svg)

### 2.1. Domain
The **Domain** contains all the business logic executed by the microservice. It consists of three main parts: **Entities**, **Use Cases**, and **Boundary**.

#### 2.1.1. Entities
**Entities** define the data model for the domain. These are simple `Go` structs used within the domain as well as across the domain boundary to transfer data.

#### 2.1.2. Usecases
**Usecases** contain all the business logic. Any external dependencies needed by the Use Cases (e.g., database resources) are injected into them using dependency inversion.

#### 2.1.3. Boundary
The **Boundary** marks the interface between the **Domain** and the **orchestration layers**. It contains contracts (`Go` interfaces) that facilitate dependency inversion.

### 2.2. Orchestration
Orchestration contains **Infrastructure**, **Presentation** and **Persistence**.

#### 2.2.1. Infrastructure
**Infrastructure** contains the lowest-level resources needed by the microservice, such as configuration and the IoC container.

#### 2.2.2. Presentation
**Presentation** contains all outward-facing interfaces. These are the communication channels between the microservice and the outside world. This is where you place your RESTful, GraphQL, gRPC, or WebSocket servers. It’s worth noting that you don’t need to implement all of these in a single microservice; it solely depends on the specifics of your implementation.

Telemetry configurations for metrics and traces can be set up here as well. However, with currently available options, I would use an [eBPF](https://ebpf.io/) collector to gather telemetry. Unless you need to export custom metrics from your service, this approach provides sufficient information about your service.

#### 2.2.3. Persistence
**Persistence** is used to hold all data-related resources, whether it’s simple file writes, an RDBMS, an object store, or even an event-sourcing system backed by a local store. The important thing to remember is that all implementation details should be encapsulated within the **Persistence** layer. The **Domain** using these resources must not know (or care) about how persistence is implemented. Saving to a static file should be no different than saving to a messaging backend from the perspective of the **Domain** layer. All complexities related to the underlying persistence technologies should remain contained within the **Persistence** layer.

## 3. Usage

**Catalyst** comes with a script to make it easy to create new projects with it. You can find this script with each release which is version locked to that specific release.

Use the following command to directly create a new microservice using **Catalyst** in your current working directory.
```shell
curl -fsSL https://github.com/kosatnkn/catalyst/releases/download/v3.0.0/new_from_v3.0.0.sh | bash -s -- --module="example.com/dummyuser/sampler"
```

If you prefer first downloading the script, inspect it and then run it (which is the safer approach), then use following commands.
```shell
# download first
curl -fsSL -o new_from_v3.0.0.sh https://github.com/kosatnkn/catalyst/releases/download/v3.0.0/new_from_v3.0.0.sh

# inspect
# ...

# once ready, run
chmod +x new_from_v3.0.0.sh
./new_from_v3.0.0.sh --module="example.com/dummyuser/sampler"
```

> **NOTE:**
>
>The directory name for your new microservice will be inferred from your Go module name which you will be passing in as the `--module` parameter.
>
> The script can handle version information in the module name when inferring a name for the directory. So both `example.com/dummyuser/sampler` and `example.com/dummyuser/sampler/v2` will produce `sampler` as the directory name.
