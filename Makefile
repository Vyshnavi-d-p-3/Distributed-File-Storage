build:
	@go build -o bin/mygoproject

run: build
	@./bin/mygoproject

test:
	@go test ./... -v
