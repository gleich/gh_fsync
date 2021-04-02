FROM golang:1.15-alpine3.12 AS builder

# Meta data
LABEL maintainer="matthewgleich@gmail.com"
LABEL description="🔄 GitHub action to sync files across repos in GitHub"

# Copying over all the files
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app .

# hadolint ignore=DL3006,DL3007
FROM alpine:latest

RUN apk update
# hadolint ignore=DL3018
RUN apk add git --no-cache

WORKDIR /
COPY --from=builder /usr/src/app/app .
