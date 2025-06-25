tidy:
	@go mod tidy

build: tidy
	@go build -o ./bin/app ./cmd/main.go

run: build
	@./bin/app

test:
	@go test -v ./...