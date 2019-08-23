all:
	make build
run:
	@go run main.go
build:
	@go build -o bin/gostman