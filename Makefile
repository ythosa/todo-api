SHELL := /bin/bash

migrate_path = ~/Programs/migrate.linux-amd64

.PHONY: build
build:
	go build -v ./cmd/main.go

.PHONY: test
test:
	go test -v -race -timeout 30s ./pkg/...

.PHONY: migrate_init
migrate_init:
	$(migrate_path) create -ext sql -dir ./schema -seq init

.PHONY: migrate_up
migrate_up:
	$(migrate_path) -path ./schema -database 'postgres://ythosa:qwerty@localhost:5432/todo?sslmode=disable' up

.PHONY: migrate_down
migrate_down:
	$(migrate_path) -path ./schema -database 'postgres://ythosa:qwerty@localhost:5432/todo?sslmode=disable' down

.DEFAULT_GOAL := build
