.PHONY: build clean kafka deploy consumer

build:
	env GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/api cmd/api/main.go
	env GOOS=linux CGO_ENABLED=1 go build -ldflags="-s -w" -o bin/consumer cmd/consumer/main.go

kafka:
	docker-compose up --build

deploy:
	./bin/api

consumer:
	./bin/consumer

clean:
	rm -rf ./bin
