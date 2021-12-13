build:
	go build -o bin/monitoring main.go

run:
	go run main.go

compile:
	GOOS=linux go build -o bin/monitoring
