GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
BINARY_NAME=messaging-server
BINARY_PATH="cmd/$(BINARY_NAME)/main.go"

# Build
build-dev:
	@echo "  > Building binary for development..."
	@$(GOBUILD) -o $(BINARY_NAME) $(BINARY_PATH)

build-prod:
	@echo "  > Building binary for production..."
	@CGO_ENABLED=0 $(GOBUILD) -a -installsuffix cgo -ldflags '-extldflags "-static"' -o $(BINARY_NAME) $(BINARY_PATH)

# Run (used for development)
run:
	@echo "  > Running server..."
	@$(GOCMD) run $(BINARY_PATH)

# Dependency download
deps:
	@echo "  > Tidying and downloading deps..."
	@$(GOMOD) tidy



