build:
	go build -o bin/monitoring main.go

run:
	go run main.go

compile:
	GOOS=linux CGO_ENABLED=0 go build -o bin/monitoring main.go
