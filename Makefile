install-tools:
	go install github.com/swaggo/swag/cmd/swag@latest
	go install github.com/golang-migrate/migrate/v4
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.57.2

docker-build:
	docker build -t trade-union-telegram-bot:latest . --file Dockerfile

build:
	go version
	go mod tidy
	CGO_ENABLED=0 go build cmd/trade-union/main.go

run:
	go mod tidy
	go run cmd/trade-union/main.go

test:
	go test -short -coverprofile=coverage.out ./...

coverage-html:
	go tool cover -html=./coverage.out

docs:
	swag init -g ./internal/app/app.go

lint:
	golangci-lint run --enable-all