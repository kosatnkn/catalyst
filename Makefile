# ref: https://bytes.usc.edu/cs104/wiki/makefile

# Testing
.PHONY: test
test:
	go test -v ./...

# Go dependency management
.PHONY: dep-upgrade-list
dep-upgrade-list:
	go list -u -m all

.PHONY: dep-upgrade-all
dep-upgrade-all:
	go get -t -u ./... && go mod tidy
