# Please DO NOT use it for build a normal Docker image for Create Go App CLI!
# This Dockerfile used ONLY with GoReleaser project (`task release [TAG...]`).

FROM alpine:edge

LABEL maintainer="Vic Shóstak <vic@shostak.dev>"

# Copy Create Go App CLI binary.
COPY cgapp /cgapp

# Install git, npm (with nodejs).
RUN apk add --no-cache git npm

# Set entry point.
ENTRYPOINT ["/cgapp"]