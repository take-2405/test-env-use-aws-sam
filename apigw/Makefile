PHONY: deps clean build

deps:
	go get -u ./...

clean:
	rm -rf ./template/template

build:
	GOOS=linux GOARCH=amd64 go build -o ./template/template ./template/main.go
