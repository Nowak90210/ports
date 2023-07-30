image:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/ports
	@docker build -t ports .

run:
	@docker run  -p 8080:8080 ports

test:
	@go test -race ./...