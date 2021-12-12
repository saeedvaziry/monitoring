build:
	go build -o bin/monitoring main.go

run:
	go run main.go

compile:
	GOOS=linux go build -ldflags "-s -w" -o bin/monitoring main.go
