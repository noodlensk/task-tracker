setup-openapi: ## Setup openapi codegen tool
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
setup-asyncap: ## Setup asyncapi codegen tool
	npm install -g @asyncapi/generator
setup-lint: ## Set up linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.49.0
setup: setup-asyncap setup-openapi setup-lint ## Setup tooling
fmt: ## gofmt and goimports all go files
	find . -name '*.go' | while read -r file; do gofumpt -w "$$file"; goimports -w "$$file"; done
lint: ## Lint
	cd internal/common && golangci-lint run
test: ## Run tests
	cd internal/common && go test -count=1 -p=8 -parallel=8 -race ./...
dep: ## Get all dependencies
	cd internal/common && go mod download && go mod tidy
start-env: ## Start the local env
	docker-compose up -d

stop-env: ## Stop the local env
	docker-compose down

restart-env: stop-env start-env ## Restart the local env

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
