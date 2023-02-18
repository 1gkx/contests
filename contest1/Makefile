# libraries versions
LINTER_VERSION=1.50.1

download.all: download.goimport download.linter

.PHONY: download.all

download.linter:
	@curl -sSfL --connect-timeout 10 https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b .bin v$(LINTER_VERSION)
	@echo "Install golint success"

.PHONY: download.linter

download.goimport:
	@go install golang.org/x/tools/cmd/goimports@v$(GOIMPORTS_VERSION)
	@echo "Install goimports success"

.PHONY: download.goimport

list:
	@grep '^[^#[:space:]].*:' Makefile

.PHONY: list

lint:
	@./.bin/golangci-lint run -v
	@echo "Done"

.PHONY: lint

test:
	go test ./...

.PHONY: test

format:
	@go fmt ./...
	@goimports -l -w ./*.go
	@echo "Done"

.PHONY: format