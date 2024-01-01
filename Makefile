build:
	@go build -o bin/main main.go
run: build
	@clear
	@go run main.go

.PHONY: build run
