# STEP 1
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV GOPATH=/go
WORKDIR $GOPATH/src/students-api

COPY . .
RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags="-w -s" -o /$GOPATH/bin/students-api cmd/server/main.go

# STEP 2
FROM scratch

WORKDIR /app
COPY --from=builder /go/bin/students-api /app/students-api

EXPOSE 9010
ENTRYPOINT ["/app/students-api"]