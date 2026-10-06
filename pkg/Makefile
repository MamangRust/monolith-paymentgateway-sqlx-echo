fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run

test:
	go test -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.out -v ./...