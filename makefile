mod:
	GO111MODULE=on go mod tidy

build:
	go build  -o cmd/emissioner main.go

test:
	go test -v ./...
