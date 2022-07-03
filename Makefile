all: clean build run

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	go mod tidy
