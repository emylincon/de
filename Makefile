format:
	go fmt ./...
	go vet ./...

test:
	go mod tidy
	go test ./...

build: format test
	go build

install: format test
	go install

clean:
	go clean -i
