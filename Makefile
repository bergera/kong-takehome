# using a Makefile to keep commands easy to remember and short

GO_ENV=GO111MODULE=on CGO_ENABLED=0

.PHONY: default
default: help

# from https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## commands help text
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: clean
clean: ## clean workspace and remove build/test artifacts
	rm -rf bin

.PHONY: build
build: clean ## compile the binary
	${GO_ENV} go build -o bin/main

.PHONY: run
run: ## launch the server locally
	@${GO_ENV} go run main.go

.PHONY: test-unit
test-unit: ## run unit tests
	@${GO_ENV} go test -v
