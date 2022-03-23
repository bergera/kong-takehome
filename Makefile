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
	docker-compose down --rmi local

.PHONY: build
build: clean ## compile the binary
	${GO_ENV} go build -o bin/main ./kong_takehome/*

.PHONY: run
run: ## launch the server locally
	@${GO_ENV} go run ./kong_takehome

.PHONY: test-unit
test-unit: ## run unit tests
	@${GO_ENV} go test -v ./kong_takehome

.PHONY: install
install: ## install npm dependencies
	@npm install

.PHONY: docker
docker: ## run docker containers
	@docker compose up

.PHONY: test-integration
test-integration: install ## run integration tests - run server separately first
	@docker compose up --wait --remove-orphans --force-recreate --build
	@npm test
