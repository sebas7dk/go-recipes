
NAME	= go-recipes
BIN 	= go-recipes

all: build

test: deps
	go test -timeout=5s -cover -race -v ./...

ci_test: test

build: deps test
	go build -v -a -o $(NAME) .

run:
	./$(BIN)

deps:
	go get -v -t ./...

fmt:
	go fmt ./...

lint:
	golint ./...

test_fast:
	go test ./...
