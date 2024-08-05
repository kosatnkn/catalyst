# Running and Testing
run:
	# update metadata
	./metadata.sh
	go run main.go

test:
	go test -v ./...

# Mocking
mock:
	docker run --init --name catalyst_mock -it --rm -v $(PWD)/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"

# Go dependancy management
dep_upgrade_list:
	go list -u -m all

dep_upgrade_all:
	go get -t -u ./... && go mod
