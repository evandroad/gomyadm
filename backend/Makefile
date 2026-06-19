.PHONY: build frontend go version run

VERSION := $(shell git describe --tags --always)

build: version frontend go

frontend:
	cd ../frontend && npm install && npm run build

go:
	swag init -g cmd/gomyadm/main.go 
	sudo cp -r ../frontend/web cmd/gomyadm
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o gomyadm ./cmd/gomyadm

run:
	swag init -g cmd/gomyadm/main.go 
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o gomyadm ./cmd/gomyadm
	./gomyadm