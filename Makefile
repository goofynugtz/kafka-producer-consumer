.PHONY: build clean deploy

build:
	env GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/start cmd/main.go

containers:
	docker-compose up --build

clean:
	rm -rf ./bin
