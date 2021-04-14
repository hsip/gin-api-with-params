.PHONY: build start test

build : 
	go build -o dist/gin_api

start :
	go run main.go

test :
	go test -v common/**