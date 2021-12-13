build:
	go build -o bin/monitoring main.go

run:
	go run main.go

compile:
	GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o bin/monitoring main.go
