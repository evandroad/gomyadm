.PHONY: build frontend go version

VERSION := $(shell git describe --tags --always)

build: version frontend go

frontend:
	cd ../frontend && npm install && npm run build

go:
	sudo cp -r ../frontend/web cmd/gomyadm
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o gomyadm ./cmd/gomyadm