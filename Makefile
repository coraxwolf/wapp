
# Version/Build Info
VERSION = $(shell git describe --tags --always --dirty)
BUILD_VERSION = $(shell git rev-parse --short HEAD)
BUILD_DATE = $(shell date '+%Y-%m-%dT%H:%M:%S')
