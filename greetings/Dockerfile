#
# Base Image & running greetings
#
FROM alpine

EXPOSE 8080

# Add git for fetching dependencies
# Ca-certificates for HTTPS
# TimeZones
# RUN apk update && apk add --no-cache git ca-certificates tzdata

WORKDIR /app
ADD ./greetings /app/

ENTRYPOINT [ "/app/greetings" ]
