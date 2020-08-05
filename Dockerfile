# This Dockerfile used ONLY with GoReleaser project (`task release [TAG...]`).
# Please DO NOT use it for build a normal Docker image for Create Go App CLI!

FROM alpine:3.12

LABEL maintainer="Vic Shóstak <truewebartisans@gmail.com>"

# Copy Create Go App CLI binary.
COPY cgapp /cgapp

# Install git, npm (with nodejs).
RUN apk add --no-cache git npm

# Install frontend CLIs (globally and in silent mode).
RUN npm i -g -s --unsafe-perm \
    create-react-app \
    preact-cli \
    @vue/cli \
    @angular/cli \
    degit

# Set entry point.
ENTRYPOINT ["/cgapp"]