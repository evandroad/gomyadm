.PHONY: build frontend go version run

VERSION := $(shell git describe --tags --always)

build: version frontend go

frontend:
	cd frontend/ && npm install && npm run build

go:
	swag init -g cmd/server/main.go 
	sudo cp -r frontend/web cmd/server
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o gomyadm ./cmd/server

run:
	swag init -g cmd/server/main.go
	go build -ldflags="-s -w -X main.Version=$(VERSION)" -o gomyadm ./cmd/server
	./gomyadm

wails:
	wails build -ldflags "-s -w -X gomyadm/internal/services.Version=$(VERSION)" -tags webkit2_41
	cp build/bin/gomyadm deb-pkg/usr/local/bin/

deb: build
	rm -f container-manager*.deb
	dpkg-deb --build deb-pkg container-manager_$(VERSION)_amd64.deb