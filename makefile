NAME := blip
BINDIR := ./build
VERSION := $(shell git describe --tags 2>/dev/null)
LDFLAGS := -X 'main.version=$(VERSION)'

.PHONY: build-darwin
build:
	@go build -o $(BINDIR)/darwin-amd64/$(NAME) -ldflags "$(LDFLAGS)" .

.PHONY: build-linux
build-linux:
	@GOOS=linux GOARCH=amd64 go build -o $(BINDIR)/linux-amd64/$(NAME) -ldflags "$(LDFLAGS)" .

.PHONY: build-windows
build-windows:
	@GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/windows-amd64/$(NAME).exe -ldflags "$(LDFLAGS)" .

.PHONY: clean
clean:
	@-rm -rf $(BINDIR)

.PHONY: help
help:
	@make2help $(MAKEFILE_LIST)
