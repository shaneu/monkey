shell := "/bin/bash"

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	go fmt ./...