cli:
	go build -o bin/book book/main.go

server:
	go build -o bin/server server/main.go

tests:
	go test ./...

run-server: 
	./bin/server

install-cli:
	go install ./book
