FROM alpine

EXPOSE 8081

# Add git for fetching dependencies
# Ca-certificates for HTTPS
# TimeZones
# RUN apk update && apk add --no-cache git ca-certificates tzdata

WORKDIR /app
ADD ./niche-todos /app/

ENTRYPOINT [ "/app/niche-todos" ]
