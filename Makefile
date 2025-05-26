run: build
	@./bin/redis --listenAddr :5832

build:
	@go build -o bin/redis ./src


test:
	@go test -timeout 30s -run ^TestNewClients$ ./client -v -count=1
