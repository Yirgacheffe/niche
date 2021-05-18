# STEP 1
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV GOPATH=/go
WORKDIR $GOPATH/src/hello

COPY . .
RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build ./cmd/main.go -ldflags="-w -s" -o /$GOPATH/bin/hello

# STEP 2
FROM scratch

COPY --from=builder /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]
