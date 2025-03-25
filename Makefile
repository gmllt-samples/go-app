APP_NAME := server
CMD_DIR := ./cmd/server
BUILD_DIR := ./bin

.PHONY: all build run clean

all: build

build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_DIR)

run: build
	@echo "Running..."
	@PORT=8080 $(BUILD_DIR)/$(APP_NAME)

clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
