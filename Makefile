.PHONY: build
build:
	go build -v ./cmd/server
	docker-compose up -d

.DEFAULT_GOAL := build