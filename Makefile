cli:
	go build -o bin/book book/main.go

server:
	go build -o bin/server server/main.go

run-server: 
	./bin/server

run-cli:
	go install ./book