# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY ./app/main /go/bin/app

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
#RUN go install golang.org/x/example/outyet

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/app