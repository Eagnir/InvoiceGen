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

build-macos:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/invoiceGen-web.macos ./interface/web/main.go
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/mac/invoiceGen-cmd.macos ./interface/cmd/main.go
	
build-windows:
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-web.exe ./interface/api/main.go
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-cmd.exe ./interface/cmd/main.go

build-api-macos: 
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o ./bin/mac/invoiceGen-web.macos ./interface/api/main.go
	
build-cmd-macos:
	env CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -a -o ./bin/mac/invoiceGen-cmd.macos ./interface/cmd/main.go
	
build-api-windows: 
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-web.exe ./interface/api/main.go

build-cmd-windows:
	env CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC=x86_64-w64-mingw32-gcc go build -a -o ./bin/windows/invoiceGen-cmd.exe ./interface/cmd/main.go

ci: dependencies test	

test:
	go test -tags testing ./...

fmt: ## gofmt and goimports all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do gofmt -w -s "$$file"; goimports -w "$$file"; done