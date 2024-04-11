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