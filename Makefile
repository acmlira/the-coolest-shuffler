all: clean build run

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	go mod tidy

test:
	go test -v -count=1 ./...

mockery-install:
	cd /tmp && go install github.com/vektra/mockery/v2@latest

mockery: mockery
	mockery --all --dir ./internal/api --output ./internal/api/mocks