PROJECT_PATH=$(shell pwd)
MODULE_NAME=sentinel
IMAGE_REPOSITORY=kjuiop
MAIN_FILE=/main.go

LDFLAGS=-X main.BUILD_TIME=`date -u '+%Y-%m-%d_%H:%M:%S'`
LDFLAGS+=-X main.GIT_HASH=`git rev-parse HEAD`
LDFLAGS+=-s -w

all: config build

config:
	@if [ ! -d $(TARGET_DIR) ]; then mkdir $(TARGET_DIR); fi

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags "$(LDFLAGS)" $(PROJECT_PATH)$(MAIN_FILE)

docker-build:
	docker build -f Dockerfile --tag $(IMAGE_REPOSITORY)/$(MODULE_NAME) .

docker-run:
	docker run --rm --name $(MODULE_NAME) --network redis-network $(IMAGE_REPOSITORY)/$(MODULE_NAME)
