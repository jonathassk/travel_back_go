build:
	@go build -o bin/goProject cmd/main.go

test:
	@go test -v ./...

run: build 
	@./bin/goProject