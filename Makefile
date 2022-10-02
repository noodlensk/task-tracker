setup-openapi: ## Setup openapi codegen tool
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
setup-asyncap: ## Setup asyncapi codegen tool
	npm install -g @asyncapi/generator
setup-lint: ## Set up linter
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.48
setup: setup-asyncap setup-openapi setup-lint ## Setup tooling
openapi_http: ## Build stubs from openapi spec for backend
	oapi-codegen --old-config-style -generate types -o internal/tasks/ports/openapi_types.gen.go -package ports api/openapi/tasks.yaml
	oapi-codegen --old-config-style -generate chi-server -o internal/tasks/ports/openapi_api.gen.go -package ports api/openapi/tasks.yaml
	oapi-codegen --old-config-style -generate types -o internal/common/clients/tasks/openapi_types.gen.go -package articles api/openapi/tasks.yaml
	oapi-codegen --old-config-style -generate client -o internal/common/clients/tasks/openapi_api.gen.go -package articles api/openapi/tasks.yaml
	oapi-codegen --old-config-style -generate types -o internal/users/ports/openapi_types.gen.go -package ports api/openapi/users.yaml
	oapi-codegen --old-config-style -generate chi-server -o internal/users/ports/openapi_api.gen.go -package ports api/openapi/users.yaml
	oapi-codegen --old-config-style -generate types -o internal/common/clients/users/openapi_types.gen.go -package articles api/openapi/users.yaml
	oapi-codegen --old-config-style -generate client -o internal/common/clients/users/openapi_api.gen.go -package articles api/openapi/users.yaml
fmt: ## gofmt and goimports all go files
	find . -name '*.go' | while read -r file; do gofumpt -w "$$file"; goimports -w "$$file"; done
dep: ## Get all dependencies
	cd internal/common && go mod download && go mod tidy
	cd internal/tasks && go mod download && go mod tidy
	cd internal/users && go mod download && go mod tidy
start-env: ## Start the local env
	docker-compose up -d

stop-env: ## Stop the local env
	docker-compose down

restart-env: stop-env start-env ## Restart the local env

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help