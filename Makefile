# ref: https://bytes.usc.edu/cs104/wiki/makefile

# Running
.PHONY: run
run:
	./metadata/set_metadata.sh $(PWD)
	go run main.go

.PHONY: run_with_env
run_with_env:
	./metadata/set_metadata.sh $(PWD)

	CATALYST_APP_NAME="env_app_name" \
	CATALYST_APP_MODE="DEBUG" \
	CATALYST_APP_PORT=8000 \
	CATALYST_APP_TIMEZONE="UTF" \
	CATALYST_APP_METRICS_ENABLED=true \
	CATALYST_APP_METRICS_PORT=8001 \
	CATALYST_DB_HOST="env_db_host" \
	CATALYST_DB_PORT=5432 \
	CATALYST_DB_DATABASE="env_db_name" \
	CATALYST_DB_USER="env_db_user" \
	CATALYST_DB_PASSWORD="env_db_pwd" \
	CATALYST_DB_POOLSIZE=5 \
	CATALYST_DB_CHECK=false \
	CATALYST_LOG_LEVEL="INFO" \
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
	go get -t -u ./... && go mod tidy
