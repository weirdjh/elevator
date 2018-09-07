include makefile_env

USERID ?= heojh93
IMAGE ?= elevator
VERSION ?= 0.1


.PHONY: build push

all: go_build image_build image_push

go_build:
	go build

image_build: Dockerfile
	docker build -t $(USERID)/$(IMAGE):$(VERSION) -f Dockerfile .

image_push:
	docker push $(USERID)/$(IMAGE):$(VERSION)