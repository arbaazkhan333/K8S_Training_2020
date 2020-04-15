.PHONY: build-image
build-image: build-binary
	docker build -t arbaazkhan/hello_k8s:v1 .
build-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello_k8s
