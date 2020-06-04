# STEP 1
FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git

ENV GOPATH=/go
WORKDIR $GOPATH/src/niche-todox

COPY . .
RUN go get -d -v

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags="-w -s" -o /$GOPATH/bin/niche-todox

# STEP 2
FROM scratch

WORKDIR /app
COPY --from=builder /go/bin/niche-todox /app/niche-todox

EXPOSE 8081
ENTRYPOINT ["/app/niche-todox"]
