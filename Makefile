all: lint test build

clean:
	rm -f ecnl
	rm -f junit.xml
	rm -f ecnl
	rm -rf output

deps:
	go mod tidy
	go install golang.org/x/tools/gopls@v0.13.2
	go install github.com/swaggo/swag/cmd/swag@v1.16.2
	go install go.uber.org/mock/mockgen@v0.2.0
	go install github.com/onsi/ginkgo/v2/ginkgo@v2.12.1
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2

swagger:
	swag init -g main.go

mocks:
	go generate -x ./...

build: clean deps swagger
	go build -o ecnl -ldflags="-s -w" main.go

test: clean mocks
	ginkgo --junit-report=junit.xml ./...

lint:
	golangci-lint run ./...

run: swagger
	go run main.go api

# Prunes both dangling images and unused containers
docker-prune:
	docker system prune --all --force

docker-build:
	docker build -t jediknights/ecnl-api .
	docker tag jediknights/ecnl-api jediknights/ecnl-api:latest

docker-build-client:
	docker build -t jediknights/ecnl-client -f ./client/Dockerfile ./client
	docker tag jediknights/ecnl-client jediknights/ecnl-client:latest

docker-push:
	docker push --all-tags jediknights/ecnl-api

docker-push-client:
	docker push --all-tags jediknights/ecnl-client

docker-mongo:
	docker run -it --rm --name mongodb -v ~/mongo/data:/data/db -p 27017:27017 mongo:latest
