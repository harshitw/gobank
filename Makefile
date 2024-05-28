# Used to bootstrap some common tasks - good practice
build:
	@go build -o bin/gobank

run: build
	@./bin/gobank

test: 
	@go test -v ./...