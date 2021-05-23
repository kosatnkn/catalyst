run:
	go run main.go

test:
	go test -v ./...

mock:
	docker run --init --name catalyst_mock -it --rm -v $(PWD)/docs/api:/tmp -p 3000:4010 stoplight/prism mock -h 0.0.0.0 "/tmp/openapi.yaml"