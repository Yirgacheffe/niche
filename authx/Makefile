BINARY_NAME=niche-auth

.PHONY: help
help:
	@printf "\033[32m\xE2\x9c\x93 usage: make [target]\n\n\033[0m"

all: test build

build:
	go build -o ${BINARY_NAME} ./cmd/main.go

run:
	go build -o ${BINARY_NAME} ./cmd/main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	go test -v ./...

vet:
	go vet ./...

deps:
	@printf "go get github.com/gorilla/mux"

image:
	docker build -t 'soline74/niche-auth:dev-4fd7739' .

# migrateup:
# migrate -path db/migration -database "postgresql://<user>:<pwd>@localhost:5432/authx?sslmode=disable" -verbose up

# migratedown:
# migrate -path db/migration -database "postgresql://<user>@<pwd>:5432/authx?sslmode=disable"           -verbose down
