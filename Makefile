run: build
	@./bin/redis

build:
	@go build -o bin/redis .
