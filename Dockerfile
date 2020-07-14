FROM alpine:3.12

LABEL maintainer="Vic Shóstak <truewebartisans@gmail.com>"

# Copy Create Go App binary.
COPY cgapp /cgapp

# Install git, npm.
RUN apk add --no-cache git npm

# Install frontend CLIs.
RUN npm i -g -s create-react-app preact-cli @vue/cli degit @angular/cli

# Set entry point.
ENTRYPOINT ["/cgapp"]