.PHONY: frontend api version run wails wails-dev deb

VERSION := $(shell git describe --tags --always)
LDFLAGS := -s -w -X gomyadm/internal/services.Version=$(VERSION)

frontend:
	cd frontend/ && npm install && npm run build -- --mode web

api: frontend
	swag init -g cmd/server/main.go 
	sudo cp -r frontend/web cmd/server
	go build -ldflags="$(LDFLAGS)" -o build/api/gomyadm ./cmd/server

run: api
	build/api/gomyadm

wails:
	wails build -ldflags="$(LDFLAGS)" -tags webkit2_41

wails-dev:
	wails dev -ldflags="$(LDFLAGS)" -tags webkit2_41

deb: version wails
	cp build/bin/gomyadm deb-pkg/usr/local/bin/
	rm -f gomyadm*.deb
	rm -rf .deb-build
	cp -a deb-pkg .deb-build
	find .deb-build -name .keep -delete
	dpkg-deb --build .deb-build gomyadm_$(VERSION)_amd64.deb
	rm -rf .deb-build