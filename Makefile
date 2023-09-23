clean:
	rm -f ecnl
	rm -f junit.xml

deps:
	go mod download
	go install github.com/onsi/ginkgo/v2/ginkgo
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

build: clean
	go build -v ./...

test: clean
	ginkgo --junit-report=junit.xml ./...

lint:
	golangci-lint run ./...
