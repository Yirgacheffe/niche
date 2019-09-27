FROM alpine

# Set arguments
ARG VCS_REF
ARG BUILD_DATE
ARG VERSION

# Build Metadata
LABEL   org.label-schema.vcs-ref=$VCS_REF \
        org.label-schema.name="Niche Web App" \
        org.label-schema.description="Simple golang web for Istio demo" \
        org.label-schema.vcs-url="https://github.com/Yirgacheffe/niche" \
        org.label-schema.build-date=$BUILD_DATE \
        org.label-schema.version=$VERSION \
        org.label-schema.docker.dockerfile="/web/Dockerfile"

ENV GIT_SHA $VCS_REF
ENV IMAGE_BUILD_DATE $BUILD_DATE
ENV APP_VERSION $VERSION

WORKDIR /app
ADD ./niche-web /app/

ENTRYPOINT /app/niche-web
EXPOSE 8080
