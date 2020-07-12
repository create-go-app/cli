FROM alpine:latest

LABEL maintainer="Vic Shóstak <truewebartisans@gmail.com>"

# Install git
RUN apk add --no-cache git npm

# Install frontend CLIs
RUN npm i -g -s create-react-app preact-cli @vue/cli degit @angular/cli

# Copy Create Go App binary
COPY dist/cgapp_linux_amd64/cgapp .

# Set entry point to /cgapp
ENTRYPOINT ["/cgapp"]