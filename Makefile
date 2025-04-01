build:
	@echo "  >  Building binaries ..."
	@goreleaser build --snapshot --clean

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