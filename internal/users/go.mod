module github.com/noodlensk/task-tracker/internal/users

go 1.19

require (
	github.com/deepmap/oapi-codegen v1.11.0
	github.com/go-chi/chi/v5 v5.0.7
	github.com/go-chi/render v1.0.2
	github.com/google/uuid v1.3.0
	github.com/noodlensk/task-tracker/internal/common v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.0
	go.uber.org/zap v1.23.0
)

require (
	github.com/ThreeDotsLabs/watermill v1.1.1 // indirect
	github.com/ajg/form v1.5.1 // indirect
	github.com/cristalhq/jwt/v4 v4.0.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-chi/cors v1.2.1 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.4 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/noodlensk/task-tracker/internal/common => ../common/
