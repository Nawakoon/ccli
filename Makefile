.PHONY: build
build:
	@echo "Building..."
	@go build -o bin/ccli cmd/main.go

.PHONY: install
install:
	@echo "installing..."
	@go build -o bin/ccli cmd/main.go
	@mkdir -p $$HOME/.ccli
	@HOME=$$(echo $$HOME); cp bin/ccli $$HOME/.local/bin/
	@echo "install done"
	@echo "run: ccli add --name <command name> --file <file>"

.PHONY: test
test:
	@echo "running unit tests..."
	@go clean -testcache && go test ./...

.PHONY: test-v
test-v:
	@echo "running unit tests verbose..."
	@go clean -testcache && go test ./... -v

.PHONY: list
list:
	@cat Makefile | grep -E '^[a-zA-Z0-9_-]+:.*'