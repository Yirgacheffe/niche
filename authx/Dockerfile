# Build
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV GOPATH=/go
WORKDIR $GOPATH/src/niche-auth

COPY . .
RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags="-w -s" -o /$GOPATH/bin/niche-auth ./cmd/main.go

# GET artifacts
FROM scratch

WORKDIR /app
COPY --from=builder /go/bin/niche-auth /app/niche-auth
COPY ./keys /app

EXPOSE 9010
ENTRYPOINT ["/app/niche-auth"]
