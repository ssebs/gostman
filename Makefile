all:
	make build
run:
	@go run main.go
build:
	@go build -o bin/gostman
doc:
	@echo "go to http://localhost:6060"
	@godoc -http=":6060"