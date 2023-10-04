clean:
	rm -f ecnl
	rm -f junit.xml
	rm -f ecnl
	rm -rf output

swagger:
	~/go/bin/swag init -g main.go

docs:

deps:
	go mod download
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build: clean
	go build -o ecnl -ldflags="-s -w" main.go

test: clean
	ginkgo --junit-report=junit.xml ./...

lint:
	golangci-lint run ./...

run: swagger
	go run main.go api

docker-mongo:
	docker run -it --rm --name mongodb -v ~/mongo/data:/data/db -p 27017:27017 mongo:latest
