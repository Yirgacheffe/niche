FROM alpine

# Set arguments
ARG VCS_REF
ARG BUILD_DATE
ARG IMAGE_TAG_REF

# Medadata
LABEL   org.label-schema.vcs-ref=$VCS_REF \
        org.label-schema.name="Niche API app" \
        org.label-schema.description="Simple golang web api for Istio demo" \
        org.label-schema.vcs-url="https://github.com/Yirgacheffe/niche" \
        org.label-schema.build-date=$BUILD_DATE \
        org.label-schema.docker.dockerfile="/api/Dockerfile"

ENV GIT_SHA ${VCS_REF}
ENV IMAGE_BUILD_DATE ${BUILD_DATE}
ENV IMAGE_TAG ${IMAGE_TAG_REF}

WORKDIR /app
ADD ./niche-api /app/

ENTRYPOINT /app/niche-api
EXPOSE 8081
