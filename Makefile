VERSION:=$(shell git describe --tags --always --dirty)
IMAGE_NAME=counter-gateway:$(VERSION)

counter-gateway: main.go go.mod go.sum
	go build -o counter-gateway

run: counter-gateway
	./counter-gateway

docker:
	docker build -t $(IMAGE_NAME) .
