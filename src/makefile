BINARY_NAME=cardhub

build:
	GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux .

run: build
	./${BINARY_NAME}-linux

clean:
	go clean
	rm ${BINARY_NAME}-linux

build-image:
	docker 