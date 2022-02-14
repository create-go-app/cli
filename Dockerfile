# Use this file for build a normal Docker/Podman image for Create Go App CLI.

FROM golang:1.17-alpine AS builder

LABEL maintainer="Vic Shóstak <vic@shostak.dev>"

# Move to working directory (/build).
WORKDIR /build

# Copy and download dependency using go mod.
COPY go.mod go.sum ./
RUN go mod download

# Copy the code into the container.
COPY . .

# Set necessary environment variables needed for our image and build the Create Go App CLI
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o cgapp cmd/cgapp/main.go

FROM alpine:edge

# Copy binary and config files from /build to root folder of scratch container.
COPY --from=builder ["/build/cgapp", "/"]

# Install git, npm (with nodejs).
RUN apk add --no-cache git npm

# Set entry point.
ENTRYPOINT ["/cgapp"]