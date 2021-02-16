.PHONY: all
all: build-prefix build-infix ## Build all targets (run `make` on its own)

.PHONY: build-prefix
build-prefix: ## Build only the prefix calculator
	@go build -o bin/prefix-calculator cmd/prefix-calculator/main.go

.PHONY: build-infix
build-infix: ## Build only the infix calculator
	@go build -o bin/infix-calculator cmd/infix-calculator/main.go

.PHONY: test
test: ## Run integration tests
	@go test -v ./tests/...

.PHONY: clean
clean: ## Remove binaries
	@rm -rf bin/

.PHONY: find_todo
find_todo: ## Find all the todo's in the comments.
	@grep --color=always --include=\*.go -PnRe '(//|/*).*TODO' --exclude-dir=.history/ ./ || true

.PHONY: help
help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'