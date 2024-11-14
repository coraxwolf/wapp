
# Version/Build Info
VERSION = $(shell git describe --tags --always --dirty)
BUILD_VERSION = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date '+%Y-%m-%dT%H:%M:%S')

# Compile Flags
LDFLAGS = -X 'main.config.Version=$(VERSION)' -X 'main.config.BuildVersion=$(BUILD_VERSION)' -X 'main.config.BuildTime=$(BUILD_TIME)'
GOFLAGS = -ldflags "$(LDFLAGS)"

# Build Name and Paths
BINARY_NAME = "wapp"
BINARY_DIR = "bin"
DEV_NAME = "wap-dev-$(BUILD_VERSION)"
DEV_DIR = "tmp"

# Temple Program
TEMPL = $(shell which templ

# Clean
clean:
	@echo "Cleaning...."
	@rm -rf $(BINARY_DIR) $(DEV_DIR)


# Build Css
css:
	@echo "Building Tailwind CSS"
	@npm run css


# TEMPL Generate
templates:
	@echo "Generating TEMPL Files"
	@$(TEMPL) generate

# Build Development
build-dev: clean css templates
	@echo "Building Development Build ($(BUILD_VERSION))"
	export ENV=development
	@go build $(GOFLAGS) -o $(DEV_DIR)/$(DEV_NAME)

# Testing

# Test
test: clean css templates
	@echo "Running Tests"
	export ENV=testing
	@go test ./...

# Coverage
cover: test
	@echo "Coverage Report"
	go test -coverprofile=coverage.out $(GOFLAGS) ./...
	@go tool cover -func=coverage.out

# Run Development
run: build-dev
	@echo "Running Development Build $(BUILD_VERSION)"
	$(DEV_DIR)/$(DEV_NAME)

# Prod Build
build: clean css templates
	@echo "Production Build..."
	export ENV=Production
	@go build $(GOFLAGS) -o $(BINARY_DIR)/$(BINARY_NAME)
	@echo "Build Complete."