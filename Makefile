all: clean test build run

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build ./cmd/main.go

run:
	go run ./cmd/main.go

clean:
	go mod tidy

test:
	go test -v -count=1 ./...

test-all:
	go test -v -count=1 -tags=all ./...

lint:
	gofmt -w .

mockery-install:
	cd /tmp && go install github.com/vektra/mockery/v2@latest

mockery: mockery-install
	mockery --all --dir ./internal/api --output ./internal/api/mocks
	mockery --all --dir ./internal/service --output ./internal/service/mocks