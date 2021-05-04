.PHONY: all
all: build
FORCE: ;

SHELL  := env LIBRARY_ENV=$(LIBRARY_ENV) $(SHELL)
LIBRARY_ENV ?= dev

BIN_DIR = $(PWD)/bin

.PHONY: build

clean:
	rm -rf bin/*

dependencies:
	go mod download

build: clean dependencies build-macos build-windows

build-macos: build-web-macos build-cli-macos
	
build-windows: build-web-windows build-cli-windows

build-web-macos: 
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/invoiceGen-web.macos ./interface/web/main.go
	
build-cli-macos:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/mac/invoiceGen-cli.macos ./interface/cli/main.go
	
build-web-windows: 
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-web.exe ./interface/web/main.go

build-cli-windows:
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-cli.exe ./interface/cli/main.go

ci: dependencies test	

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done