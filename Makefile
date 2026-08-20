.PHONY: frontend swagger api version run wails wails-dev deb cli

VERSION := $(shell git describe --tags --always)
LDFLAGS := -s -w -X gomyadm/internal/services.Version=$(VERSION)

all: api deb cli

frontend:
	cd frontend/ && bun install && bun run build -- --mode web

swagger:
	swag init -g cmd/server/main.go

run: swagger
	go run -tags dev ./cmd/server

api: frontend
	rm -r cmd/server/web 
	cp -r frontend/web cmd/server
	go build -ldflags="$(LDFLAGS)" -o build/api/gomyadm ./cmd/server

wails:
	wails build -ldflags="$(LDFLAGS)" -tags webkit2_41

wails-dev:
	wails dev -ldflags="$(LDFLAGS)" -tags webkit2_41

cli:
	go build -ldflags="$(LDFLAGS)" -o build/cli/gomyadm ./cmd/cli

version:
	sed -i 's/^Version:.*/Version: $(VERSION)/' deb-pkg/DEBIAN/control

deb: version wails
	cp build/bin/gomyadm deb-pkg/usr/local/bin/
	rm -f gomyadm*.deb
	rm -rf .deb-build
	cp -a deb-pkg .deb-build
	find .deb-build -name .keep -delete
	dpkg-deb --build .deb-build gomyadm_$(VERSION)_amd64.deb
	rm -rf .deb-build