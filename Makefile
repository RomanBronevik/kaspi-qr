.DEFAULT_GOAL := build

BINARY_NAME = svc
BUILD_PATH = cmd/build

fmt_swdoc:
	swag fmt

swdoc:
	swag init --parseDependency --parseDepth 3

build:
	mkdir -p $(BUILD_PATH)
	CGO_ENABLED=0 go build -o $(BUILD_PATH)/$(BINARY_NAME) main.go

clean:
	rm -rf $(BUILD_PATH)
