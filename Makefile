# BUILD_FLAGS=-X 'main.GitCommitHash=`git rev-parse --short HEAD`' -X 'main.BuiltAt=`date +%FT%T%z`'
BUILD_FLAGS=-X 'main.BuiltAt=`date +%FT%T%z`'
build:
	@echo "  >  Building binary ..."
	@go build -o fkr -ldflags="$(BUILD_FLAGS)" ./cmd/fedramp-ksi-report

test:
	@echo "  >  Validating code ..."
	@go vet ./...
	@go clean -testcache
	@go test ./...

tidy:
	@echo "  >  Tidying go.mod ..."
	@go mod tidy

lint:
	@echo "  >  Linting code ..."
	@golangci-lint run --timeout 5m

local-release:
	@echo "  >  Releasing ..."
	@goreleaser build --single-target --snapshot --clean