.PHONY: build test clean run

BINARY_NAME=gbot

build:
	go build -o $(BINARY_NAME) -v

test:
	go test -v ./...

clean:
	go clean
	rm -f $(BINARY_NAME)

run:
	go run main.go

deps:
	go mod download

lint:
	golangci-lint run
