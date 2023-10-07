all: lint test build

clean:
	rm -f ecnl
	rm -f junit.xml
	rm -f ecnl
	rm -rf output

deps:
	go mod tidy
	go install github.com/swaggo/swag/cmd/swag@v1.16.2
	go install go.uber.org/mock/mockgen@v0.2.0
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.12.1
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

swagger:
	swag init -g main.go

mocks:
	go generate -v ./...

build: clean deps swagger
	go build -o ecnl -ldflags="-s -w" main.go

test: clean mocks
	ginkgo --junit-report=junit.xml ./...

lint:
	golangci-lint run ./...

run: swagger
	go run main.go api

docker-build:
	docker build -t ecnl-api .

docker-tag:
	docker tag ecnl-api jediknights/ecnl-api:latest
	docker tag ecnl-api jediknights/ecnl-api:$(date +%s)

docker-push:
	docker push --all-tags jediknights/ecnl-api

docker-mongo:
	docker run -it --rm --name mongodb -v ~/mongo/data:/data/db -p 27017:27017 mongo:latest
