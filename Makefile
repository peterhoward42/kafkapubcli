
GOLANGCI_LINT_DEP=github.com/golangci/golangci-lint/cmd/golangci-lint

.PHONY: run
run: test
run:
	go run cmd/kafkapubcli.go

.PHONY: test
test: style
test:
	go test ./...

.PHONY: style
style: fmt
	golangci-lint run --enable=gofmt -v

.PHONY: fmt
fmt: build
	gofmt -w -s -l .

.PHONY: build
build: deps
	go build  ./...

.PHONY: deps
deps: ## Install dependencies
	go get $(GOLANGCI_LINT_DEP)


