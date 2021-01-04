.PHONY: build
build:
	go build -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./pkg/...

.DEFAULT_GOAL := build
