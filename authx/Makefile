BINARY_NAME=niche-auth

.PHONY: help
help:
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"

all: test build

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test -v main.go

deps:
	@printf "go get github.com/gorilla/mux"
