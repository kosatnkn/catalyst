# ref: https://bytes.usc.edu/cs104/wiki/makefile

# Running
.PHONY: run
run:
	./metadata/set_metadata.sh $(PWD)
	go run main.go

.PHONY: run_in_docker
run_in_docker: docker_build
	docker run --name catalyst-test-api --rm -p 8000:8000 -p 8001:8001 kosatnkn/catalyst-test-api:latest

# Testing
.PHONY: test
test:
	go test -v ./...

# Spin up a mock API using the OpenAPI document
.PHONY: mock
mock:
	docker run --init --name catalyst_mock -it --rm -v $(PWD)/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"

# Containerizing
.PHONY: docker_build
docker_build:
	./scripts/set_metadata.sh
	docker build --tag kosatnkn/catalyst-test-api:latest .

# Go dependency management
.PHONY: dep_upgrade_list
dep_upgrade_list:
	go list -u -m all

.PHONY: dep_upgrade_all
dep_upgrade_all:
	go get -t -u ./... && go mod
