.PHONY: build-prefix
build-prefix:
	@go build -o bin/prefix-calculator ./cmd/prefix-calculator/main.go

.PHONY: test
test:
	@go test -v ./tests/...

.PHONY: all
all: build-prefix

.PHONY: clean
clean:
	@rm -rf bin/