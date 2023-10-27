.PHONY: build clean deploy

build:
	env GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/start main.go

run-consumer:
	@go run cmd/consumer/main.go

clean:
	rm -rf ./bin

