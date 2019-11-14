# Build ref:
# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
#

#
# Build binary
#
FROM golang:1.13-alpine AS builder

EXPOSE 8000

# Add git for fetching dependencies
# Ca-certificates for HTTPS
# TimeZones
RUN apk update && apk add --no-cache git ca-certificates tzdata

# Create appuser.
RUN adduser -D -g '' appuser

# Set ENV variables
ENV GOPATH /go
# ENV GOPROXY https://goproxy.io

# Copy files
WORKDIR $GOPATH/src/greetings
COPY main.go .

# Go get dependencies.
RUN go get -d -v

# Run go mod if you have (temprary comment it)
# RUN go mod download
# RUN go mod verify

# Set go environment variables
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# Build the binary
RUN go build -ldflags="-w -s" -a -installsuffix cgo -o $GOPATH/bin/greetings .

#
# Build the image
#
FROM scratch

# COPY fils from builder
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs
COPY --from=builder /etc/passwd /etc/passwd

# COPY binary file
COPY --from=builder $GOPATH/bin/greetings /greetings

# Use 'appuser'
USER appuser

# Run greetings binary application
ENTRYPOINT ["/greetings"]
