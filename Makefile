build:
	go build -o bin/monitoring main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o bin/monitoring main.go
