
SERVICE_NAME ?= $(shell basename $(dir $(abspath $(firstword $(MAKEFILE_LIST)))))
APP_NAME     = backend

DOCKERFILE   = docker/default/Dockerfile
DOCKER_IMAGE = loghole/$(SERVICE_NAME)

VERSION  ?= $$(git describe --tags)
GIT_HASH := $$(git rev-parse HEAD)

build: 
	docker build \
	--build-arg APP_NAME=$(APP_NAME) \
	--build-arg SERVICE_NAME=$(SERVICE_NAME) \
	--build-arg GIT_HASH=$(GIT_HASH) \
	--build-arg VERSION=$(VERSION) \
	-f $(DOCKERFILE) \
	-t $(DOCKER_IMAGE) \
	-t $(DOCKER_IMAGE):$(VERSION) \
	.

default-image:
	docker build -t loghole/$(APP_NAME):latest -f docker/default/Dockerfile .

build-frontend:
	go-bindata -fs -o backend/bindata/bindata.go -prefix "./frontend/dist" -pkg "bindata" ./frontend/dist/...