FROM alpine

# Set arguments
ARG VCS_REF
ARG BUILD_DATE
ARG VERSION

ENV GIT_SHA $VCS_REF
ENV IMAGE_BUILD_DATE $BUILD_DATE
ENV APP_VERSION $VERSION

WORKDIR /app
ADD ./niche-web /app/
ADD ./tmpls/index.html /app/tmpls/

ENTRYPOINT ["/app/niche-web"]
EXPOSE 8082
