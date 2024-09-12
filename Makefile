BUILD_DIR=./bin
MAIN_DIR=./cmd

BINARY_NAME=static-site-generator

run:
	go run $(MAIN_DIR)

build:
	rm -rf $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_DIR)

launch:
	$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
	go clean

test:
	go clean -testcache && go test ./...

.PHONY: run build launch clean test