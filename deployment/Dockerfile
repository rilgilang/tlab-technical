# STEP 1: Build the Go binary
FROM golang:1.24.9-alpine3.22 AS builder

ARG GO_BUILD_COMMAND="go build -tags static_all"

# Install some build deps + ssh tools for the setup below.
RUN apk update && apk --no-cache add  build-base  git bash  coreutils openssh  openssl


# this command if you get source from bitbucket repos
# Create the directory where the application will reside
RUN mkdir -p /go/src/gus-payment-be


WORKDIR /go/src/gus-payment-be

COPY . .


# application builder step
RUN go mod tidy && go mod download && go mod vendor
RUN eval $GO_BUILD_COMMAND


# STEP 2 build a small image
# Set up the final (deployable/runtime) image.
FROM alpine:3.17.0


# setup package dependencies
RUN apk --no-cache update

ENV BUILDDIR=/go/src/gus-payment-be
ENV PROJECT_DIR=/opt/gus-payment-be

# Setting timezone
ENV TZ=Asia/Jakarta
RUN apk add -U tzdata
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Set ENV variable for Datadog Build Version
ARG BUILD_VERSION
ENV DD_VERSION=$BUILD_VERSION

#create project directory
RUN mkdir -p $PROJECT_DIR/config

WORKDIR $PROJECT_DIR

COPY --from=builder $BUILDDIR/gus-payment-be gus-payment-be


CMD ["sh","-c", "/opt/gus-payment-be/gus-payment-be"]
