all:
	go build -v ./...
	go test -v -cover ./...
