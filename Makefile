# ref: https://bytes.usc.edu/cs104/wiki/makefile

# Build
.PHONY: build
build:
	@./metadata/set_metadata.sh $(pwd) \
	&& go build -v -o main .

# Running
.PHONY: run
run:
	@./metadata/set_metadata.sh $(pwd) \
	&& go run .

.PHONY: run-with-env
run-with-env:
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
	go run .

.PHONY: run-in-docker
run-in-docker: docker-build
	docker run --name catalyst-test-api --rm -p 8000:8000 -p 8001:8001 kosatnkn/catalyst-test-api:latest

# Testing
.PHONY: test
test:
	go test -covermode=count -coverpkg=./... -coverprofile=cover.out ./...

# Spin up a mock API using the OpenAPI document
.PHONY: mock
mock:
	docker run --init --name catalyst_mock -it --rm -v $(PWD)/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"

# Containerizing
.PHONY: docker-build
docker-build:
	./scripts/set_metadata.sh
	docker build --tag kosatnkn/catalyst-test-api:latest .

# Go dependency management
.PHONY: dep-upgrade-list
dep-upgrade-list:
	go list -u -m all

.PHONY: dep-upgrade-all
dep-upgrade-all:
	go get -t -u ./... && go mod tidy

SHELL := /bin/bash
.PHONY: release
release:
	@echo "Checking if on 'main' branch ..."; \
	git fetch origin main >/dev/null 2>&1; \
	CURRENT_BRANCH=$$(git rev-parse --abbrev-ref HEAD); \
	if [ "$$CURRENT_BRANCH" != "main" ]; then \
		echo "Error: Not on 'main' branch (currently on: $$CURRENT_BRANCH)"; \
		exit 1; \
	fi; \
	echo "Checking if 'main' has local uncommitted changes..."; \
	if ! git diff --quiet || ! git diff --cached --quiet; then \
    echo "Error: Uncommitted changes exist in 'main'"; \
    git status -s; \
    exit 1; \
	fi; \
	echo "Checking if local 'main' branch is up to date with 'origin/main'..."; \
	LOCAL_HASH=$$(git rev-parse main); \
	REMOTE_HASH=$$(git rev-parse origin/main); \
	if [ "$$LOCAL_HASH" != "$$REMOTE_HASH" ]; then \
		echo "Error: 'main' is not up to date with 'origin/main'"; \
		exit 1; \
	fi; \
	read -p "Enter new release tag (vX.Y.Z): " TAG; \
	REGEX="^v(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)\\.(0|[1-9][0-9]*)$$"; \
	if [[ ! "$$TAG" =~ $$REGEX ]]; then \
		echo "Error: Invalid version format for '$$TAG'"; \
		echo "Version must follow vX.Y.Z format (e.g., v1.0.0, v2.1.5)."; \
		exit 1; \
	fi; \
	echo "Creating new tag $$TAG and drafting a new release..."; \
	git tag --annotate $$TAG --message="Release $$TAG" && git push origin $$TAG
