# ref: https://bytes.usc.edu/cs104/wiki/makefile
.PHONY: run test mock docker_run docker_build dep_upgrade_list dep_upgrade_all

# Running
run:
	./metadata/set_metadata.sh $(PWD)
	go run main.go

run_in_docker: docker_build
	docker run --name catalyst-test-api --rm -p 8000:8000 -p 8001:8001 kosatnkn/catalyst-test-api:latest

# Testing
test:
	go test -v ./...

# Spin up a mock API using the OpenAPI document
mock:
	docker run --init --name catalyst_mock -it --rm -v $(PWD)/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"

# Containerizing
docker_build:
	./scripts/set_metadata.sh
	docker build --tag kosatnkn/catalyst-test-api:latest .

# Go dependency management
dep_upgrade_list:
	go list -u -m all

dep_upgrade_all:
	go get -t -u ./... && go mod
