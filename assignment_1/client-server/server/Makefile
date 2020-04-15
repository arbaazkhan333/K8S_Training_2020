.PHONY: build-image
build-image: build-binary
	docker build -t arbaazkhan/server:v1 .
build-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server
