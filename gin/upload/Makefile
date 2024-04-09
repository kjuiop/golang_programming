PROJECT_PATH=$(shell pwd)
MODULE_NAME=upload-app

BUILD_NUM_FILE=build.num
BUILD_NUM=$$(cat ./build.num)
APP_VERSION=$$(cat ./version.txt)
TARGET_VERSION=$(APP_VERSION).$(BUILD_NUM)
IMAGE_REPOSITORY="kjuiop"

TARGET_DIR=bin
OUTPUT=$(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)
MAIN_DIR=/main.go
LDFLAGS=-X main.BUILD_TIME=`date -u '+%Y-%m-%d_%H:%M:%S'`
LDFLAGS+=-X main.GIT_HASH=`git rev-parse HEAD`
LDFLAGS+=-s -w

all: config target-version docker-build docker-push

config:
	@if [ ! -d $(TARGET_DIR) ]; then mkdir $(TARGET_DIR); fi

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" -o $(OUTPUT) $(PROJECT_PATH)$(MAIN_DIR)
	cp $(OUTPUT) ./$(MODULE_NAME)

docker-build:
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY):$(TARGET_VERSION) .

docker-push:
	docker push $(IMAGE_REPOSITORY):$(TARGET_VERSION)

docker-release:
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY):latest .
	docker push $(IMAGE_REPOSITORY):latest

target-version:
	@echo "========================================"
	@echo "APP_VERSION    : $(APP_VERSION)"
	@echo "BUILD_NUM      : $(BUILD_NUM)"
	@echo "TARGET_VERSION : $(TARGET_VERSION)"
	@echo "========================================"

build-num:
	@echo $$(($$(cat $(BUILD_NUM_FILE)) + 1 )) > $(BUILD_NUM_FILE)

clean:
	rm -f $(PROJECT_PATH)/$(TARGET_DIR)/$(MODULE_NAME)*

